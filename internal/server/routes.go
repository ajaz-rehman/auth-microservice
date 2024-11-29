package server

import "net/http"

type Routes map[string]http.HandlerFunc

var signupHandler = requestHandler(signup)

func getRoutes() Routes {
	routes := Routes{
		"POST /signup": signupHandler,
	}

	return routes
}
