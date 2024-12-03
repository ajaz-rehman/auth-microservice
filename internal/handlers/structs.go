package handlers

import (
	"net/http"

	"github.com/ajaz-rehman/auth-microservice/internal/app"
)

type Handler func(*app.App) http.HandlerFunc
