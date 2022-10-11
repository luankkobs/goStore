package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectionDatabase() *sql.DB {
	connection := "user=postgres dbname=alura_store password=vacapreta1 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
