package main

import (
	_ "embed"
	"log"

	_ "github.com/lib/pq"
	"github.com/olyop/objectgraph-go/schema/database"
)

//go:embed schema.sql
var schema string

func main() {
	loadEnv()

	database.Connect()
	defer database.Close()

	_, err := database.DB.Exec(schema)
	if err != nil {
		log.Default().Fatal(err)
	}
}
