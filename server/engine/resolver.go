package engine

import (
	"context"

	"github.com/olyop/graphql-go/server/engine/distributedcache"
	"github.com/olyop/graphql-go/server/engine/inmemorycache"
)

func Resolver[R any](ctx context.Context, options ResolverOptions) (*R, error) {
	result, err := resolve[*R](ctx, options)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return result, nil
}

func ResolverList[R any](ctx context.Context, options ResolverOptions) ([]*R, error) {
	result, err := resolve[*[]*R](ctx, options)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return *result, nil
}

func resolve[R any](ctx context.Context, options ResolverOptions) (R, error) {
	cacheKey := getCacheKey(options)

	// first check the in-memory cache
	result, found := inmemorycache.Get[R](options.RetrieverKey, cacheKey)
	if found {
		return result, nil
	}

	// this lock is to ensure that only one resolver is fetching the value from the source
	requestMutex := getRequestMutext(ctx, cacheKey)
	requestMutex.Lock()
	defer requestMutex.Unlock()

	// check again in case another resolver has already fetched the value
	result, found = inmemorycache.Get[R](options.RetrieverKey, cacheKey)
	if found {
		return result, nil
	}

	cacheDuration := getCacheDuration(options)
	retriever := getRetriever(options)

	// check the distributed cache
	result, found, err := distributedcache.Get[R](ctx, cacheKey)
	if err != nil {
		return result, err
	}

	// if found, cache the result in the in-memory cache and return the value
	if found {
		inmemorycache.Set(options.RetrieverKey, cacheKey, result, cacheDuration)

		return result, nil
	}

	// this lock is to prevent multiple requests from retrieving the same value from the source
	resolverMutex := getResolverMutex(ctx, cacheKey)
	resolverMutex.Lock()
	defer resolverMutex.Unlock()

	// check again in case another request has already fetched and cached the result
	result, found = inmemorycache.Get[R](options.RetrieverKey, cacheKey)
	if found {
		return result, nil
	}

	// not found in any of the caches, so retrieve the result
	data, err := retriever(ctx, options.RetrieverArgs)
	if err != nil {
		return result, err
	}

	result = data.(R)

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
