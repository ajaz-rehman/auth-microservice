package main

import (
	"log/slog"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/server"
	"github.com/ajaz-rehman/auth-microservice/internal/utils"
)

func main() {
	config, err := utils.LoadConfig()

	if err != nil {
		slog.Error("utils.LoadConfig: " + err.Error())
		os.Exit(1)
	}

	slog.Info("Starting server on port: " + config.Env.PORT)

	err = server.ListenAndServe(config.Env.PORT)

	if err != nil {
		slog.Error("server.ListenAndServe: " + err.Error())
		os.Exit(1)
	}
}
