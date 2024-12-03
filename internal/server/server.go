package server

import (
	"log/slog"
	"net/http"

	"github.com/ajaz-rehman/auth-microservice/internal/app"
)

func ListenAndServe() error {
	app, err := app.GetApp()

	if err != nil {
		return err
	}

	slog.Info("Starting server on port: " + app.ENV.PORT)

	mux := getMuxWithRoutes(app)

	server := &http.Server{
		Addr:    ":" + app.ENV.PORT,
		Handler: mux,
	}

	return server.ListenAndServe()
}

func getMuxWithRoutes(app *app.App) *http.ServeMux {
	mux := http.NewServeMux()
	routes := GetRoutes()

	for _, route := range routes {
		mux.HandleFunc(route.Pattern, route.Handler(app))
	}

	return mux
}
