package routes

import "net/http"

type Route struct {
	Pattern string
	Handler http.HandlerFunc
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
