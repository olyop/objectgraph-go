package objectgraph

import (
	"errors"
	"io/fs"

	"github.com/graphql-go/graphql"
	"github.com/redis/go-redis/v9"
)

type Configuration struct {
	SourceFiles fs.FS
	CachePrefix string
	CacheRedis  *redis.Options
	Objects     ConfigurationObjects
	Scalars     map[string]*graphql.Scalar
}

type ConfigurationObjects map[string]*ConfigurationObject

type ConfigurationObject struct {
	PrimaryKey string
	Retrievers any
}

func (c *Configuration) validate() error {
	if c.CacheRedis == nil {
		return errors.New("cacheRedis is required")
	}

	if len(c.Objects) == 0 {
		return errors.New("objects is required")
	}

	for key := range c.Objects {
		if key == "" {
			return errors.New("an object key is required")
		}

		if key == "Query" {
			return errors.New("query is a reserved key")
		}

		if c.Objects[key] == nil {
			return errors.New("object config is required")
		}

		if c.Objects[key].PrimaryKey == "" {
			return errors.New("PrimaryKey is required")
		}

		if c.Objects[key].Retrievers == nil {
			return errors.New("retrievers is required")
		}
	}

	return nil
}

func (c *Configuration) getTypeNames() []string {
	types := make([]string, 0, len(c.Objects)+1)

	for key := range c.Objects {
		types = append(types, key)
	}

	types = append(types, "Query")

	return types
}
