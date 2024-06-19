package inmemorycache

import "time"

func Set[T any](groupKey string, cacheKey string, value *T, ttl time.Duration) {
	groupCache := handleGroup(groupKey)

	groupCache.Set(cacheKey, value, ttl)
}
