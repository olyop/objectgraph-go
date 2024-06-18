package graphqlops

import (
	"context"

	"github.com/olyop/graphql-go/server/graphqlops/distributedcache"
	"github.com/olyop/graphql-go/server/graphqlops/inmemorycache"
)

func Resolver[R any](ctx context.Context, options ResolverOptions) (*R, error) {
	e := ctx.Value(EngineContextKey{}).(*Engine)

	result, err := e.resolve(ctx, options)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return result.(*R), nil
}

func ResolverList[R any](ctx context.Context, options ResolverOptions) ([]*R, error) {
	e := ctx.Value(EngineContextKey{}).(*Engine)

	result, err := e.resolve(ctx, options)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return result.([]*R), nil
}

func (e *Engine) resolve(ctx context.Context, options ResolverOptions) (any, error) {
	cacheKey := e.getCacheKey(options)
	cacheDuration := e.getCacheDuration(options)
	retriever := e.getRetriever(options)

	// first check the in-memory cache
	result, found := inmemorycache.Get(options.RetrieverKey, cacheKey)
	if found {
		return result, nil
	}

	// this lock is to ensure that only one resolver is fetching the value from the source
	requestResolverLocker := loadResolverLocker(ctx, ResolverRequestLockerContextKey{}, cacheKey)
	requestResolverLocker.Lock()
	defer requestResolverLocker.Unlock()

	// check again in case another resolver has already fetched the value
	result, found = inmemorycache.Get(options.RetrieverKey, cacheKey)
	if found {
		return result, nil
	}

	// check the distributed cache
	result, found, err := distributedcache.Get(ctx, cacheKey)
	if err != nil {
		return result, err
	}

	// if found, cache the result in the in-memory cache and return the value
	if found {
		inmemorycache.Set(options.RetrieverKey, cacheKey, result, cacheDuration)

		return result, nil
	}

	// this lock is to prevent multiple requests from retrieving the same value from the source
	resolverLocker := loadResolverLocker(ctx, ResolverLockerContextKey{}, cacheKey)
	resolverLocker.Lock()
	defer resolverLocker.Unlock()

	// check again in case another request has already fetched and cached the result
	result, found = inmemorycache.Get(options.RetrieverKey, cacheKey)
	if found {
		return result, nil
	}

	// not found in any of the caches, so retrieve the result
	result, err = retriever(ctx, options.RetrieverArgs)
	if err != nil {
		return result, err
	}

	// cache the result in the in-memory and distributed caches
	inmemorycache.Set(options.RetrieverKey, cacheKey, result, cacheDuration)
	err = distributedcache.Set(ctx, cacheKey, result, cacheDuration)
	if err != nil {
		inmemorycache.Delete(options.RetrieverKey, cacheKey)
		return result, err
	}

	return result, nil
}

type ResolverOptions struct {
	CacheDuration string
	RetrieverKey  string
	RetrieverArgs RetrieverArgs
}
