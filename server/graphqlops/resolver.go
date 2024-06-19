package graphqlops

import (
	"context"

	"github.com/olyop/graphqlops-go/graphqlops/distributedcache"
	"github.com/olyop/graphqlops-go/graphqlops/inmemorycache"
)

func Resolver[R any](ctx context.Context, options ResolverOptions) (*R, error) {
	e := ctx.Value(EngineContextKey{}).(*Engine)

	result, err := resolve[R](e, ctx, options)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return result, nil
}

func ResolverList[R any](ctx context.Context, options ResolverOptions) ([]*R, error) {
	e := ctx.Value(EngineContextKey{}).(*Engine)

	result, err := resolve[[]*R](e, ctx, options)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return *result, nil
}

func resolve[T any](e *Engine, ctx context.Context, options ResolverOptions) (*T, error) {
	cacheKey := e.getCacheKey(options)
	cacheDuration := e.getCacheDuration(options)
	retriever := e.getRetriever(options)

	// first check the in-memory cache
	result, found := inmemorycache.Get[T](options.RetrieverKey, cacheKey)
	if found {
		return result, nil
	}

	// this lock is to ensure that only one resolver is fetching the value from the source
	requestResolverLocker := e.loadResolverLocker(ctx, ResolverRequestLockerContextKey{}, cacheKey)
	requestResolverLocker.Lock()
	defer requestResolverLocker.Unlock()

	// check again in case another resolver has already fetched the value
	result, found = inmemorycache.Get[T](options.RetrieverKey, cacheKey)
	if found {
		return result, nil
	}

	// check the distributed cache
	result, found, err := distributedcache.Get[T](ctx, cacheKey)
	if err != nil {
		return result, err
	}

	// if found, cache the result in the in-memory cache and return the value
	if found {
		inmemorycache.Set(options.RetrieverKey, cacheKey, result, cacheDuration)

		return result, nil
	}

	// this lock is to prevent multiple requests from retrieving the same value from the source
	resolverLocker := e.loadResolverLocker(ctx, ResolverLockerContextKey{}, cacheKey)
	resolverLocker.Lock()
	defer resolverLocker.Unlock()

	// check again in case another request has already fetched and cached the result
	result, found = inmemorycache.Get[T](options.RetrieverKey, cacheKey)
	if found {
		return result, nil
	}

	// not found in any of the caches, so retrieve the result
	data, err := retriever(ctx, options.RetrieverArgs)
	if err != nil {
		return result, err
	}

	result = data.(*T)

	// set the inmemorycache here so other resolvers/requests can use it
	inmemorycache.Set(options.RetrieverKey, cacheKey, result, cacheDuration)

	// set the distributed cache here so other instances can use it
	// and in a go routine so the resolver can return the result to the client immediately
	go func() {
		err := distributedcache.Set(ctx, cacheKey, result, cacheDuration)
		if err != nil {
			inmemorycache.Delete(options.RetrieverKey, cacheKey)
		}
	}()

	return result, nil
}

type ResolverOptions struct {
	CacheDuration string
	RetrieverKey  string
	RetrieverArgs RetrieverArgs
}
