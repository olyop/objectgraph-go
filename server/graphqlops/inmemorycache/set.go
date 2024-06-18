package inmemorycache

import "time"

func Set(groupKey string, cacheKey string, value any, ttl time.Duration) {
	groupCache := handleGroup(groupKey)

	groupCache.Set(cacheKey, value, ttl)
}
