package objectgraph

import (
	"time"

	"github.com/google/uuid"
	"github.com/olyop/objectgraph/objectgraph/internal/objectcache"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func NewEngine(configuration *Configuration) (*Engine, error) {
	if err := configuration.validate(); err != nil {
		return nil, err
	}

	schemaSources := []*ast.Source{{Name: "schema.graphql", Input: configuration.Schema}}
	schemaAst, graphqlErr := gqlparser.LoadSchema(schemaSources...)
	if graphqlErr != nil {
		return nil, graphqlErr
	}

	retreivers, err := createRetrievers(configuration)
	if err != nil {
		return nil, err
	}

	schema, err := compileSchema(schemaAst, configuration, retreivers)
	if err != nil {
		return nil, err
	}

	objectcache, err := objectcache.New(configuration.Cache.Prefix, configuration.Cache.Redis)
	if err != nil {
		return nil, err
	}

	objectcache.Set(
		"Product",
		"5d227690-5d0f-4a74-8cb1-05141ad480e5", "object",
		map[string]any{
			"productID": uuid.MustParse("5d227690-5d0f-4a74-8cb1-05141ad480e5"),
			"name":      "Product 1",
		},
		time.Hour,
	)

	objectcache.Get(
		"Product",
		"5d227690-5d0f-4a74-8cb1-05141ad480e5",
		"object",
		time.Hour,
		retreivers["Product"].typ,
	)

	e := &Engine{
		schemaAst:   schemaAst,
		schema:      schema,
		retreivers:  retreivers,
		objectcache: objectcache,
	}

	return e, nil
}
