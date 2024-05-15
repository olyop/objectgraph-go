package main

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/olyop/graphql-go/database"
	"github.com/olyop/graphql-go/resolvers"
	"github.com/olyop/graphql-go/schema"

	_ "github.com/lib/pq"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

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
