package app

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func loadDB(url string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", url)

	return
}
