package parser

import (
	"io/fs"

	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func LoadSchema(schemaFs fs.FS) (*ast.Schema, error) {
	schemaSources, err := readSchema(schemaFs)
	if err != nil {
		return nil, err
	}

	schema, graphQLErr := gqlparser.LoadSchema(schemaSources...)
	if graphQLErr != nil {
		return nil, graphQLErr
	}

	return schema, nil
}
