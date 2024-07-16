package objectgraph

import (
	"github.com/graphql-go/graphql"
	"github.com/olyop/objectgraph/objectgraph/internal/objectcache"
	"github.com/olyop/objectgraph/objectgraph/internal/parser"
)

type Engine struct {
	config      *Configuration
	schema      *graphql.Schema
	objectCache *objectcache.ObjectCache
}

func NewEngine(config *Configuration) (*Engine, error) {
	_, err := parser.LoadSchema(config.SourceFiles)
	if err != nil {
		return nil, err
	}

	err = config.validate()
	if err != nil {
		return nil, err
	}

	typeNames := config.getTypeNames()

	objectcache, err := objectcache.New(config.CachePrefix, config.CacheRedis, typeNames)
	if err != nil {
		return nil, err
	}

	schemaConfig, err := compileSchemaConfig(config, objectcache)
	if err != nil {
		return nil, err
	}

	schema, err := graphql.NewSchema(*schemaConfig)
	if err != nil {
		return nil, err
	}

	e := &Engine{
		config:      config,
		schema:      &schema,
		objectCache: objectcache,
	}

	return e, nil
}

func (e *Engine) Close() {
	e.objectCache.Close()
}
