package server

import "net/http"

type Route struct {
	Pattern string
	Handler http.HandlerFunc
}
