package main

import (
	"log/slog"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/core"
	"github.com/ajaz-rehman/auth-microservice/internal/server"
)

func main() {
	config, err := core.LoadConfig()

	if err != nil {
		slog.Error("core.LoadConfig: " + err.Error())
		os.Exit(1)
	}

	slog.Info("Starting server on port: " + config.Env.PORT)

	err = server.ListenAndServe(config.Env.PORT)

	if err != nil {
		slog.Error("server.ListenAndServe: " + err.Error())
		os.Exit(1)
	}
}
