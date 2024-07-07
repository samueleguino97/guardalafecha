package main

import (
	"guardalafecha/internal/server"
	"log"
)

func main() {
	app, cleanup := server.NewServer(3000)
	defer cleanup()
	log.Println("Server started on port 3000")
	app.ListenAndServe()
}
