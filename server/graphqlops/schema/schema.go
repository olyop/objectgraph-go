package schema

import (
	"io/fs"
	"log"

	"github.com/graph-gophers/graphql-go"
)

type Schema struct {
	*graphql.Schema
}

func New(schemaFs fs.FS, resolver interface{}) *Schema {
	return &Schema{
		Schema: Parse(schemaFs, resolver),
	}
}

func Parse(schemaFs fs.FS, resolver interface{}) *graphql.Schema {
	schemaString, err := readSchema(schemaFs)
	if err != nil {
		log.Fatal(err)
	}

	options := []graphql.SchemaOpt{
		graphql.MaxParallelism(500),
		graphql.UseFieldResolvers(),
	}

	schema, err := graphql.ParseSchema(schemaString, resolver, options...)
	if err != nil {
		log.Fatal(err)
	}

	return schema
}
