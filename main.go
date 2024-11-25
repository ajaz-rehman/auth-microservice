package main

import (
	"log/slog"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/server"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	mux := server.GetMuxWithRoutes()

	slog.Info("Starting server on port " + port)

	err := server.ListenAndServe(port, mux)

	if err != nil {
		slog.Error(err.Error())
	}
}
