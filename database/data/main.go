package main

import (
	_ "github.com/lib/pq"
	"github.com/olyop/objectgraph-go/data/database"
	"github.com/olyop/objectgraph-go/data/files"
	"github.com/olyop/objectgraph-go/data/populate"
)

func main() {
	loadEnv()

	data := files.Read()

	database.Connect()
	defer database.Close()

	populate.Execute(data)
}
