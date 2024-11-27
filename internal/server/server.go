package server

import "net/http"

func getMuxWithRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	endpoints := getEndpoints()

	for _, endpoint := range endpoints {
		mux.HandleFunc(endpoint.Pattern, endpoint.Handler)
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
