package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", s.HealthHandler)
	r.Post("/", s.InvitationsHandler)

	return r
}

func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) InvitationsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) StatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
