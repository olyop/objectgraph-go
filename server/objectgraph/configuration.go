package objectgraph

import (
	"io/fs"

	"github.com/olyop/objectgraph/objectgraph/objectcache"
)

type EngineConfiguration struct {
	SourceFiles fs.FS
	Retrievers  RetrieverMap
	Cache       *objectcache.Configuration
}

func (c *EngineConfiguration) getTypes() []string {
	types := make([]string, 0, len(c.Retrievers))

	for key := range c.Retrievers {
		types = append(types, key)
	}

	return types
}

func (c *EngineConfiguration) validate() {
	if c.Cache.Redis == nil {
		panic("Cache.Redis is required")
	}

	if c.Cache.Prefix == "" {
		panic("Cache.Prefix is required")
	}

	for key := range c.Retrievers {
		if key == "" {
			panic("Retrievers key is required")
		}

		if c.Retrievers[key] == nil {
			panic("Retrievers value is required")
		}
	}
}
