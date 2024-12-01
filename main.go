package main

import (
	"log/slog"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		slog.Error("PORT is required")
	}

	slog.Info("Starting server on port: " + port)

	err := server.ListenAndServe(port)

	if err != nil {
		slog.Error("server.ListenAndServe: " + err.Error())
		os.Exit(1)
	}
}
