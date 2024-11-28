package server

import "net/http"

type Routes map[string]http.HandlerFunc

func getRoutes() Routes {
	routes := Routes{
		"POST /v1/auth/signup": requestHandler(signupHandler),
	}

	return routes
}
