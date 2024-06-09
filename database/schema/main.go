package main

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/olyop/graphql-go/schema/database"
)

func main() {
	loadEnv()

	database.Connect()
	defer database.Close()

	schemaFile, err := os.ReadFile("schema.sql")
	if err != nil {
		panic(err)
	}

	schema := string(schemaFile)

	_, err = database.DB.Exec(schema)
	if err != nil {
		panic(err)
	}
}
