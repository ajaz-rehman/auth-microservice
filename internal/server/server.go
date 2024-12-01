package server

import (
	"net/http"

	"github.com/ajaz-rehman/auth-microservice/internal/routes"
)

func getMuxWithRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	routes := routes.GetRoutes()

	for pattern, handler := range routes {
		mux.HandleFunc(pattern, handler)
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
