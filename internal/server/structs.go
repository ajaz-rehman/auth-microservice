package server

import (
	"github.com/ajaz-rehman/auth-microservice/internal/handlers"
)

type Route struct {
	Pattern string
	Handler handlers.Handler
}
