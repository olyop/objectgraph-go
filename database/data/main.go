package main

import (
	_ "github.com/lib/pq"
	"github.com/olyop/graphql-go/data/database"
	"github.com/olyop/graphql-go/data/files"
	"github.com/olyop/graphql-go/data/populate"
)

func main() {
	loadEnv()

	data := files.Read()

	database.Connect()
	defer database.Close()

	populate.Execute(data)
}
