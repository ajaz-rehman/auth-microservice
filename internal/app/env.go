package app

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func loadEnvironment() (env Environment, err error) {
	godotenv.Load()

	requiredEnvVars := []string{"PORT", "GO_ENV", "JWT_SECRET", "DATABASE_URL"}

	for _, envVar := range requiredEnvVars {
		_, exists := os.LookupEnv(envVar)

		if !exists {
			err = errors.New("Missing required environment variable: " + envVar)
			return
		}
	}

	env.PORT = os.Getenv("PORT")
	env.GO_ENV = os.Getenv("GO_ENV")
	env.JWTSecret = os.Getenv("JWT_SECRET")
	env.DATABASE_URL = os.Getenv("DATABASE_URL")

	return
}
