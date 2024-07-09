package database

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var err error

	hostname := os.Getenv("POSTGRES_HOSTNAME")
	database := os.Getenv("POSTGRES_DATABASE")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	ctimeout := os.Getenv("POSTGRES_CTIMEOUT")

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s connect_timeout=%s sslmode=disable", hostname, username, password, database, ctimeout)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(runtime.GOMAXPROCS(0))

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	db.Close()
}
