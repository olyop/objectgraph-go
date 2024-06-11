package schema

import (
	"io/fs"
	"log"

	"github.com/graph-gophers/graphql-go"
)

func Parse(schemaFs fs.FS, resolver interface{}) *graphql.Schema {
	schemaString, err := readSchema(schemaFs)
	if err != nil {
		log.Fatal(err)
	}

	options := []graphql.SchemaOpt{
		graphql.MaxParallelism(500),
		graphql.UseFieldResolvers(),
	}

	return graphql.MustParseSchema(schemaString, resolver, options...)
}
