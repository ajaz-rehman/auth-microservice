package main

import (
	"log/slog"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	logger := slog.Default()

	if port == "" {
		port = "5000"
	}

	mux := server.GetMuxWithRoutes()

	logger.Info("Starting server on port " + port)

	err := server.ListenAndServe(port, mux)

	if err != nil {
		logger.Error(err.Error())
	}
}
