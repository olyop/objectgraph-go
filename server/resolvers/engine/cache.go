package engine

import (
	"time"

	"github.com/olyop/graphql-go/server/cache"
)

func Resolver[T interface{}](options ResolverOptions[T]) (T, error) {
	cacheResult, found := cache.Get[T](options.GroupKey, options.CacheKey)
	if found {
		return cacheResult, nil
	}

	result, err := options.Retrieve()
	if err != nil {
		return result, err
	}

	cache.Set(options.GroupKey, options.CacheKey, result, options.Duration)

	return result, nil
}

type ResolverOptions[T interface{}] struct {
	GroupKey string
	Duration time.Duration
	CacheKey string
	Retrieve func() (T, error)
}
