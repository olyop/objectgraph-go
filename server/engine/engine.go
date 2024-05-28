package engine

import (
	"fmt"
	"sync"
	"time"

	"github.com/olyop/graphql-go/server/engine/cache"
)

func Resolver[R any](options ResolverOptions[R]) (R, error) {
	mu := getResolverMutext(fmt.Sprintf("%s-%s", options.GroupKey, options.CacheKey))

	mu.Lock()
	defer mu.Unlock()

	result, found := cache.Get[R](options.GroupKey, options.CacheKey)
	if found {
		return result, nil
	}

	result, err := options.Retrieve()
	if err != nil {
		return result, err
	}

	cache.Set(options.GroupKey, options.CacheKey, result, options.Duration)

	return result, nil
}

var resolverMutexMap = new(sync.Map)

func getResolverMutext(key string) *sync.Mutex {
	mu, found := resolverMutexMap.Load(key)
	if !found {
		mu = &sync.Mutex{}
		resolverMutexMap.Store(key, mu)
	}

	return mu.(*sync.Mutex)
}

type ResolverOptions[T any] struct {
	GroupKey string
	Duration time.Duration
	CacheKey string
	Retrieve func() (T, error)
}
