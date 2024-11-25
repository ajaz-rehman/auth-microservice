package main

import (
	"log/slog"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/server"
	"github.com/ajaz-rehman/auth-microservice/internal/utils"
)

func main() {
	logger := slog.Default()
	env, err := utils.LoadEnv()

	if err != nil {
		logger.Error("Error loading environment variables: " + err.Error())
		os.Exit(1)
	}

	mux := server.GetMuxWithRoutes()

	logger.Info("Starting server on port: " + env.PORT)

	err = server.ListenAndServe(env.PORT, mux)

	if err != nil {
		logger.Error("Error starting server: " + err.Error())
		os.Exit(1)
	}
}
