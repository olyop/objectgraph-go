package graphqlops

import (
	"sync"

	"github.com/olyop/graphql-go/server/graphqlops/distributedcache"
)

func NewEngine(config *Configuration) (*Engine, error) {
	err := distributedcache.Connect(config.Cache.Prefix, config.Cache.Redis)
	if err != nil {
		return nil, err
	}

	schema, err := parseSchema(config.Schema, config.Resolvers)
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
