package server

import "github.com/ajaz-rehman/auth-microservice/internal/handlers"

func GetRoutes() []Route {
	return []Route{
		{
			Pattern: "POST /signup",
			Handler: handlers.SignupHandler,
		},
	}
}
