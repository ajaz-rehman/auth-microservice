package routes

import (
	"github.com/ajaz-rehman/auth-microservice/internal/controllers"
)

func GetRoutes() []Route {
	return []Route{
		{
			Pattern:    "POST /signup",
			Controller: controllers.Signup,
		},
	}
}
