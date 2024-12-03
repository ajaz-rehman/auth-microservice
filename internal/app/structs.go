package app

import "database/sql"

type Environment struct {
	PORT         string
	GO_ENV       string
	JWTSecret    string
	DATABASE_URL string
}

type App struct {
	DB  *sql.DB
	ENV Environment
}
