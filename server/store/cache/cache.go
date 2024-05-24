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

	item, exists := group.values[cacheKey]
	if !exists {
		return value, false
	}

	value = item.value.(T)

	// expired cache item
	if item.expires.Before(time.Now()) {
		group.mu.Lock()

		// check again in case another goroutine updated the item
		item, exists = group.values[cacheKey]
		if !exists {
			// the cache item was deleted by another goroutine
			group.mu.Unlock()

			return value, false
		}

		value = item.value.(T)

		delete(group.values, cacheKey)

		group.mu.Unlock()

		return value, false
	}

	return value, true
}

func GetList[T any](groupKey string, cacheKey string) ([]*T, bool) {
	group := handleGroup(groupKey)

	var value []*T

	item, exists := group.values[cacheKey]
	if !exists {
		return value, false
	}

	value = item.value.([]*T)

	// expired cache item
	if item.expires.Before(time.Now()) {
		group.mu.Lock()

		// check again in case another goroutine updated the item
		item, exists = group.values[cacheKey]
		if !exists {
			// the cache item was deleted by another goroutine
			group.mu.Unlock()

			return value, false
		}

		value = item.value.([]*T)

		delete(group.values, cacheKey)

		group.mu.Unlock()

		return value, false
	}

	return value, true
}

func Set(groupKey string, cacheKey string, value map[string]any, ttl time.Duration) {
	group := handleGroup(groupKey)

	group.mu.Lock()

	group.values[cacheKey] = &CacheItem{
		value:   value,
		expires: time.Now().Add(ttl),
	}

	group.mu.Unlock()
}

func Exists(groupKey string, cacheKey string) bool {
	group := handleGroup(groupKey)

	item, exists := group.values[cacheKey]
	if !exists {
		return false
	}

	if item.expires.Before(time.Now()) {
		group.mu.Lock()

		delete(group.values, cacheKey)

		group.mu.Unlock()

		return false
	}

	return true
}

func handleGroup(groupKey string) *CacheGroup {
	group, initialized := cache.groups[groupKey]

	if !initialized {
		cache.mu.Lock()

		// Check again in case another goroutine initialized the group
		group, initialized = cache.groups[groupKey]
		if initialized {
			cache.mu.Unlock()
			return group
		}

		group = &CacheGroup{
			mu:     sync.Mutex{},
			values: make(CacheValues),
		}

		cache.groups[groupKey] = group

		cache.mu.Unlock()
	}

	return group
}

type Cache struct {
	mu     sync.Mutex
	groups CacheGroups
}

type CacheGroups map[string]*CacheGroup

type CacheGroup struct {
	mu     sync.Mutex
	values CacheValues
}

type CacheValues map[string]*CacheItem

type CacheItem struct {
	value   any
	expires time.Time
}
