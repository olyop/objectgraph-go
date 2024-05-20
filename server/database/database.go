package database

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func Connect() error {
	var err error

	connectionString := "user=postgres password=postgres dbname=ollies-bottleo sslmode=disable"

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	return nil
}

func Close() {
	db.Close()
}
