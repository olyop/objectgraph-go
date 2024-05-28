package cache

import (
	"sync"
	"time"
)

var cache Cache = Cache{
	mu:     sync.Mutex{},
	groups: make(CacheGroups),
}

func Get[T any](groupKey string, cacheKey string) (T, bool) {
	group := handleGroup(groupKey)

	var value T

	mapItem, exists := group.Load(cacheKey)
	if !exists {
		return value, false
	}

	item := mapItem.(*CacheItem)
	value = item.value.(T)

	if item.expires.Before(time.Now()) {
		// expired cache item
		group.Delete(cacheKey)

		return value, false
	}

	return value, true
}

func Set(groupKey string, cacheKey string, value any, ttl time.Duration) {
	group := handleGroup(groupKey)

	group.Store(cacheKey, &CacheItem{
		value:   value,
		expires: time.Now().Add(ttl),
	})
}

func Exists(groupKey string, cacheKey string) bool {
	group := handleGroup(groupKey)

	mapItem, exists := group.Load(cacheKey)
	if !exists {
		return false
	}

	item := mapItem.(*CacheItem)

	if item.expires.Before(time.Now()) {
		// expired cache item
		group.Delete(cacheKey)

		return false
	}

	return true
}

func handleGroup(groupKey string) *sync.Map {
	group, initialized := cache.groups[groupKey]

	if !initialized {
		cache.mu.Lock()

		// Check again in case another goroutine initialized the group
		group, initialized = cache.groups[groupKey]
		if initialized {
			cache.mu.Unlock()
			return group
		}

		group = new(sync.Map)

		cache.groups[groupKey] = group

		cache.mu.Unlock()
	}

	return group
}

type Cache struct {
	mu     sync.Mutex
	groups CacheGroups
}

type CacheGroups map[string]*sync.Map

type CacheItem struct {
	value   any
	expires time.Time
}
