package app

import (
	"database/sql"

	"github.com/ajaz-rehman/auth-microservice/internal/database"
	_ "github.com/lib/pq"
)

func loadDB(url string) (db *database.Queries, err error) {
	postgres, err := sql.Open("postgres", url)

	if err != nil {
		return
	}

	db = database.New(postgres)

	return
}
