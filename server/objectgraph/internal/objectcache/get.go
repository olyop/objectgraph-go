package objectcache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

func (oc *ObjectCache) Get(
	groupKey string,
	cacheKey string,
	ttl time.Duration,
) (any, bool, error) {
	object, exists := oc.inmemoryGet(groupKey, cacheKey)
	if exists {
		return object, true, nil
	}

	// not in the cache, lock the object as only one goroutine should be fetching it
	objectLocker := oc.objectLocker(groupKey, cacheKey)
	objectLocker.Lock()
	defer objectLocker.Unlock()

	// check again for where another goroutine has already fetched the object
	object, exists = oc.inmemoryGet(groupKey, cacheKey)
	if exists {
		return object, true, nil
	}

	// fetch the object from redis
	object, exists, err := oc.redisGet(groupKey, cacheKey)
	if err != nil {
		return nil, false, err
	}

	if exists {
		oc.inmemorySet(groupKey, cacheKey, object, ttl)

		return object, true, nil
	}

	return nil, false, nil
}

func (oc *ObjectCache) inmemoryGet(groupKey string, cacheKey string) (any, bool) {
	objectCache := oc.objectCache[groupKey]

	object, exists := objectCache.Get(cacheKey)
	if exists {
		return object, true
	}

	return nil, false
}

func (oc *ObjectCache) redisGet(groupKey string, cacheKey string) (any, bool, error) {
	redisKey := oc.redisKey(groupKey, cacheKey)

	data, err := oc.redis.Get(context.Background(), redisKey).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	result := make(map[string]any)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, false, err
	}

	return result, true, nil
}
