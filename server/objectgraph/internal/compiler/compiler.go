package compiler

import (
	"github.com/graphql-go/graphql"
	"github.com/olyop/objectgraph/objectgraph/objectcache"
	"github.com/vektah/gqlparser/ast"
)

func CompileSchema(ast *ast.Schema, objectcache *objectcache.ObjectCache) (*graphql.Schema, error) {
	c := &compiler{
		ast:         ast,
		objectcache: objectcache,
	}

	schemaConfig := c.compile()

	schema, err := graphql.NewSchema(*schemaConfig)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

type compiler struct {
	ast         *ast.Schema
	objectcache *objectcache.ObjectCache
}
