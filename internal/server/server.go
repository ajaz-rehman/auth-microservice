package server

import (
	"net/http"
)

func getEndpoints() Endpoints {
	endpoints := Endpoints{
		"POST /signup": signupHandler,
	}

	return endpoints
}

func getMuxWithRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	endpoints := getEndpoints()

	for path, handler := range endpoints {
		mux.HandleFunc(path, handler)
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
