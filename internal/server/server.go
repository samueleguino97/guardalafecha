package server

import (
	"database/sql"
	"fmt"
	"guardalafecha/internal/db"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/tursodatabase/go-libsql"
)

type Server struct {
	Port    int
	Db      *sql.DB
	Queries *db.Queries
}

func NewDb() (*sql.DB, func()) {
	dbName := "local.db"
	primaryUrl := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
		libsql.WithAuthToken(authToken),
		libsql.WithSyncInterval(time.Minute),
	)
	if err != nil {
		fmt.Println("Error creating connector:", err)
		os.Exit(1)
	}

	db := sql.OpenDB(connector)
	cleanup := func() {
		db.Close()
		connector.Close()
	}

	return db, cleanup
}

func NewServer(port int) (*http.Server, func()) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	sqliteDb, cleanup := NewDb()

	queries := db.New(sqliteDb)
	newServer := Server{
		Port:    port,
		Db:      sqliteDb,
		Queries: queries,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.Port),
		Handler:      newServer.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return server, cleanup

}
func (s *Server) ExtractSubdomain(r *http.Request) string {
	parts := strings.Split(r.Host, ".")
	return parts[0]
}
