package server

import "net/http"

type Endpoints map[string]http.HandlerFunc

func getEndpoints() Endpoints {
	endpoints := Endpoints{
		"POST /v1/auth/signup": requestHandler(signupHandler),
	}

	return endpoints
}
