package graphqlops

import (
	"context"
	"io/fs"
	"sync"
	"time"

	"github.com/olyop/graphqlops-go/graphqlops/graphql"
	"github.com/redis/go-redis/v9"
)

type EngineContextKey struct{}
type ResolverLockerContextKey struct{}
type ResolverRequestLockerContextKey struct{}

type Engine struct {
	configuration  *Configuration
	schema         *graphql.Schema
	resolverLocker *sync.Map
}

type Configuration struct {
	Schema         fs.FS
	Resolvers      any
	Retrievers     any
	CacheDurations CacheDurationMap
	CachePrefix    string
	CacheRedis     *redis.Options
}

type CacheDurationMap map[string]time.Duration
type Retriever func(ctx context.Context, args RetrieverArgs) (any, error)
type RetrieverArgs map[string]any
