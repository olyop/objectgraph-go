package engine

import (
	"log"
	"sort"
	"time"
)

const (
	emptyCacheKey = "none"
)

func getCacheDuration(options ResolverOptions) time.Duration {
	cacheDuration, cacheDurationFound := cacheDurations[options.CacheDuration]
	if !cacheDurationFound {
		log.Fatal("cache duration not found")
	}

	return cacheDuration
}

func getRetriever(options ResolverOptions) Retriever {
	retriever, retrieverFound := retrievers[options.RetrieverKey]
	if !retrieverFound {
		log.Fatalf("retriever not found: %s\n", options.RetrieverKey)
	}

	return retriever
}

func getCacheKey(options ResolverOptions) string {
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
		sorted = append(sorted, [2]string{key, value})
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
