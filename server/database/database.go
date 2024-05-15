package database

import "database/sql"

var db *sql.DB

func Connect() (err error) {
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/ollies-bottleo?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db.Ping()
}

func Close() {
	db.Close()
}
