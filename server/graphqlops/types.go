package graphqlops

import (
	"context"
	"io/fs"
	"sync"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/redis/go-redis/v9"
)

type Engine struct {
	schema         *graphql.Schema
	configuration  *Configuration
	resolverLocker *sync.Map
}

type Configuration struct {
	Schema     fs.FS
	Resolvers  any
	Retrievers any
	Cache      *CacheConfiguration
}

type CacheConfiguration struct {
	Durations CacheDurationMap
	Prefix    string
	Redis     *redis.Options
}

type CacheDurationMap map[string]time.Duration

type DistributedCacheOptions struct {
	URL      string
	Password string
	Prefix   string
}

type EngineContextKey struct{}
type ResolverRequestLockerContextKey struct{}
type ResolverLockerContextKey struct{}

type Retriever func(ctx context.Context, args RetrieverArgs) (any, error)
type RetrieverArgs map[string]any
