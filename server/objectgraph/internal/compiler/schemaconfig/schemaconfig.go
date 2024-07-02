package schemaconfig

import (
	"github.com/graphql-go/graphql"
	"github.com/olyop/objectgraph/objectgraph/configuration"
	"github.com/olyop/objectgraph/objectgraph/objectcache"
	"github.com/vektah/gqlparser/ast"
)

func Parse(ast *ast.Schema, objectcache *objectcache.ObjectCache, config *configuration.Configuration) graphql.SchemaConfig {
	parser := &schemaConfigParser{
		ast:         ast,
		objectcache: objectcache,
		config:      config,
	}

	// _ = schemaCompileTypes(ast)

	return graphql.SchemaConfig{
		Query: parser.query(),
	}
}

type schemaConfigParser struct {
	ast         *ast.Schema
	objectcache *objectcache.ObjectCache
	config      *configuration.Configuration
}
