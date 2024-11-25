package main

import (
	"log"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/server"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	mux := server.GetMuxWithRoutes()
	err := server.ListenAndServe(port, mux)

	if err != nil {
		log.Fatalf("server.ListenAndServe: %v", err)
	}

	log.Printf("Server listening on port %s", port)
}
