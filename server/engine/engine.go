package engine

import (
	"github.com/olyop/graphql-go/server/engine/cache"
	"time"
)

func Resolver[R any](options ResolverOptions[R]) (*R, error) {
	cacheResult, found := cache.Get[R](options.GroupKey, options.CacheKey)
	if found {
		return cacheResult, nil
	}

	result, err := options.Retrieve()
	if err != nil {
		return nil, err
	}

	cache.Set(options.GroupKey, options.CacheKey, result, options.Duration)

	return result, nil
}

func Resolvers[R any](options ResolversOptions[R]) ([]*R, error) {
	cacheResult, found := cache.GetList[R](options.GroupKey, options.CacheKey)
	if found {
		return cacheResult, nil
	}

	result, err := options.Retrieve()
	if err != nil {
		return nil, err
	}

	cache.Set(options.GroupKey, options.CacheKey, result, options.Duration)

	return result, nil
}

type ResolverOptions[T any] struct {
	GroupKey string
	Duration time.Duration
	CacheKey string
	Retrieve func() (*T, error)
}

type ResolversOptions[T any] struct {
	GroupKey string
	Duration time.Duration
	CacheKey string
	Retrieve func() ([]*T, error)
}
