package server

import (
	"net/http"

	"github.com/ajaz-rehman/auth-microservice/internal/app"
)

type Route struct {
	Pattern string
	Handler Handler
}

type Handler func(*app.App) http.HandlerFunc
