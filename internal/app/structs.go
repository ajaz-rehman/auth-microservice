package app

import (
	"github.com/ajaz-rehman/auth-microservice/internal/database"
)

type Environment struct {
	PORT         string
	GO_ENV       string
	JWTSecret    string
	DATABASE_URL string
}

type App struct {
	DB  *database.Queries
	ENV Environment
}
