package schema

import (
	"log"

	"github.com/graph-gophers/graphql-go"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
)

func Parse() *graphql.Schema {
	schemaString, err := readSchema()
	if err != nil {
		log.Fatal(err)
	}

	options := []graphql.SchemaOpt{
		graphql.MaxParallelism(20),
		graphql.UseFieldResolvers(),
	}

	return graphql.MustParseSchema(schemaString, &resolvers.Resolver{}, options...)
}
