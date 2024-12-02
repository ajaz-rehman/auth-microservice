package routes

import (
	"github.com/ajaz-rehman/auth-microservice/internal/controllers"
)

type Route struct {
	Pattern    string
	Controller controllers.Controller
}
