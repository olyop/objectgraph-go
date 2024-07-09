package parser

import (
	// "encoding/json"
	"io/fs"
	// "os"

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

	// json, err := json.MarshalIndent(schema, "", "  ")
	// if err != nil {
	// 	return nil, err
	// }

	// os.WriteFile("/home/op/code/objectgraph-go/schema.json", []byte(json), 0644)

	return schema, nil
}
