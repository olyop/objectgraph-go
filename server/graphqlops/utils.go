package graphqlops

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sort"
	"time"
)

const (
	emptyCacheKey = "none"
)

func (e *Engine) getCacheDuration(options ResolverOptions) time.Duration {
	cacheDuration, cacheDurationFound := e.configuration.Cache.Durations[options.CacheDuration]
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

	for _, arg := range sortMapAlphabetically(options.RetrieverArgs) {
		cacheKey += concatCacheKey(arg[0], arg[1])
	}

	if cacheKey == "" {
		return concatCacheKey(options.RetrieverKey, emptyCacheKey)
	}

	return concatCacheKey(options.RetrieverKey, cacheKey)
}

func sortMapAlphabetically(m RetrieverArgs) [][2]string {
	sorted := make([][2]string, 0, len(m))

	for key, value := range m {
		sorted = append(sorted, [2]string{key, value.(fmt.Stringer).String()})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
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
