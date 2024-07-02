package objectgraph

import (
	"github.com/graphql-go/graphql"
	"github.com/olyop/objectgraph/objectgraph/internal/compiler"
	"github.com/olyop/objectgraph/objectgraph/internal/parser"
	"github.com/olyop/objectgraph/objectgraph/objectcache"
)

type Engine struct {
	objectCache *objectcache.ObjectCache
	config      *EngineConfiguration
	schema      *graphql.Schema
}

func NewEngine(config *EngineConfiguration) (*Engine, error) {
	config.validate()

	types := config.getTypes()

	schemaAst, err := parser.LoadSchema(config.SourceFiles)
	if err != nil {
		return nil, err
	}

	objectcache, err := objectcache.New(config.Cache)
	if err != nil {
		return nil, err
	}

	schema, err := compiler.CompileSchema(schemaAst, objectcache)
	if err != nil {
		return nil, err
	}

	e := &Engine{
		schema:      schema,
		config:      config,
		objectCache: objectcache,
	}

	return e, nil
}

func (e *Engine) Close() {
	e.objectCache.Close()
}
