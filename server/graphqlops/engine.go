package graphqlops

import (
	"sync"

	"github.com/olyop/graphqlops-go/graphqlops/distributedcache"
	"github.com/olyop/graphqlops-go/graphqlops/parser"
)

func NewEngine(config *Configuration) (*Engine, error) {
	validateConfig(config)

	err := distributedcache.Connect(config.CachePrefix, config.CacheRedis)
	if err != nil {
		return nil, err
	}

	schema, err := parser.Exec(config.Schema, config.Resolvers)
	if err != nil {
		return nil, err
	}

	return &Engine{
		schema:         schema,
		configuration:  config,
		resolverLocker: new(sync.Map),
	}, nil
}

func (*Engine) Close() {
	distributedcache.Close()
}

func validateConfig(config *Configuration) {
	if config.CacheRedis == nil {
		panic("CacheRedis is required")
	}

	if config.CachePrefix == "" {
		panic("CachePrefix is required")
	}
}
