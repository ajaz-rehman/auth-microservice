package routes

import "net/http"

type Routes map[string]http.HandlerFunc

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
