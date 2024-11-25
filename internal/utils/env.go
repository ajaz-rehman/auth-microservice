package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (Env, error) {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		return Env{}, errors.New("PORT is required")
	}

	envType := ENV_TYPE(os.Getenv("GO_ENV"))

	if envType != DEVELOPMENT && envType != PRODUCTION && envType != TEST {
		return Env{}, errors.New("invalid GO_ENV value" + string(envType))
	}

	env := Env{
		PORT:   port,
		GO_ENV: envType,
	}

	return env, nil
}
