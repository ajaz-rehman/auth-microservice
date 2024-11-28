package server

import "net/http"

func getMuxWithRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	routes := getRoutes()

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
