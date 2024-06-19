package graphqlops

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sort"
	"sync"
	"time"
)

const (
	emptyCacheKey = "none"
)

func (e *Engine) loadResolverLocker(ctx context.Context, contextKey any, cacheKey string) *sync.Mutex {
	m := ctx.Value(contextKey).(*sync.Map)

	mu, found := m.Load(cacheKey)
	if !found {
		mu = &sync.Mutex{}
		m.Store(cacheKey, mu)
	}

	return mu.(*sync.Mutex)
}

func (e *Engine) getCacheDuration(options ResolverOptions) time.Duration {
	cacheDuration, cacheDurationFound := e.configuration.CacheDurations[options.CacheDuration]
	if !cacheDurationFound {
		log.Fatal("cache duration not found")
	}

	return cacheDuration
}

func (e *Engine) getRetriever(options ResolverOptions) Retriever {
	key := options.RetrieverKey

	// get the retriever from e.Configuration.Retrievers using reflect
	retriever := reflect.ValueOf(e.configuration.Retrievers).MethodByName(key)
	if !retriever.IsValid() {
		log.Fatalf("retriever not found: %s", key)
	}

	return retriever.Interface().(func(ctx context.Context, args RetrieverArgs) (any, error))
}

func (e *Engine) getCacheKey(options ResolverOptions) string {
	var cacheKey string

	for _, arg := range transformArgs(options.RetrieverArgs) {
		cacheKey += concatCacheKey(arg[0], arg[1])
	}

	if cacheKey == "" {
		return concatCacheKey(options.RetrieverKey, emptyCacheKey)
	}

	return concatCacheKey(options.RetrieverKey, cacheKey)
}

func transformArgs(m RetrieverArgs) [][2]string {
	sorted := make([][2]string, 0, len(m))

	for key, value := range m {
		sorted = append(sorted, [2]string{key, value.(fmt.Stringer).String()})
	}

	sort.SliceStable(sorted, func(i, j int) bool {
		key1 := sorted[i][0]
		key2 := sorted[j][0]

		// sort alphabetically by key
		return key1 < key2
	})

	return sorted
}

func concatCacheKey(values ...string) string {
	var cacheKey string

	for i, value := range values {
		if i > 0 {
			cacheKey += ":"
		}

		cacheKey += value
	}

	return cacheKey
}
