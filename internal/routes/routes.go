package routes

import (
	"github.com/ajaz-rehman/auth-microservice/internal/controllers"
)

func GetRoutes() Routes {
	Routes := Routes{
		"POST /signup": requestHandler(controllers.Signup),
	}

	return Routes
}
