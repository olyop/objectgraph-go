package main

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/resolvers"
	"github.com/olyop/graphql-go/server/schema"

	_ "github.com/lib/pq"
)

func main() {
	loadEnv()

	database.Connect()
	defer database.Close()

	schemaString, err := schema.ReadSourceFiles("./schema")
	if err != nil {
		log.Fatal(err)
	}

	schema := graphql.MustParseSchema(schemaString, &resolvers.Resolver{})

	http.Handle("/graphql", &relay.Handler{Schema: schema})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
