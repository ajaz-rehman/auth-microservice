package server

import (
	"net/http"

	"github.com/ajaz-rehman/auth-microservice/internal/routes"
)

func getMuxWithRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	routes := routes.GetRoutes()

	for _, route := range routes {
		mux.HandleFunc(route.Pattern, route.Handler)
	}

	return mux
}

func ListenAndServe(port string) error {
	mux := getMuxWithRoutes()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return server.ListenAndServe()
}
