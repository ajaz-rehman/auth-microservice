package server

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/routes"
	"github.com/joho/godotenv"
)

func getMuxWithRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	routes := routes.GetRoutes()

	for _, route := range routes {
		mux.HandleFunc(route.Pattern, route.Handler)
	}

	return mux
}

func ListenAndServe() error {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		return errors.New("PORT environment variable not set")
	}

	slog.Info("Starting server on port: " + port)

	mux := getMuxWithRoutes()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return server.ListenAndServe()
}
