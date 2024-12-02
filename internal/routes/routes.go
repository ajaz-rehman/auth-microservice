package routes

import (
	"github.com/ajaz-rehman/auth-microservice/internal/controllers"
)

func GetRoutes() []Route {
	return []Route{
		{
			Pattern: "POST /signup",
			Handler: requestHandler(controllers.Signup),
		},
	}
}
