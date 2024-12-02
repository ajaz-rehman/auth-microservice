package main

import (
	"log/slog"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/server"
)

func main() {
	err := server.ListenAndServe()

	if err != nil {
		slog.Error("server.ListenAndServe: " + err.Error())
		os.Exit(1)
	}
}
