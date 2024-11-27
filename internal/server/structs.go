package server

import "net/http"

type Endpoint struct {
	Pattern     string
	Handler     http.HandlerFunc
	RequestBody map[string]interface{}
}

type Endpoints []Endpoint
