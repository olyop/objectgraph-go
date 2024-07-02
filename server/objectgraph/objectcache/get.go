package objectcache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func (oc *ObjectCache) Get(
	groupKey string,
	cacheKey string,
	ttl time.Duration,
) (any, bool, error) {
	cacheGroup := oc.getCacheGroup(groupKey)

	object, exists := cacheGroup.Get(cacheKey)
	if exists {
		return object, true, nil
	}

	// not in the cache, lock the object as only one goroutine should be fetching it
	objectlocker := oc.getObjectLocker(groupKey, cacheKey)
	objectlocker.Lock()
	defer objectlocker.Unlock()

	// check again for where another goroutine has already fetched the object
	object, exists = cacheGroup.Get(cacheKey)
	if exists {
		return object, true, nil
	}

	// fetch the object from redis
	object, exists, err := oc.redisGet(groupKey, cacheKey)
	if err != nil {
		return "", false, err
	}

	if exists {
		cacheGroup.Set(cacheKey, object, ttl)

		return object, true, nil
	}

	return "", false, nil
}

func (oc *ObjectCache) redisGet(groupKey string, cacheKey string) (string, bool, error) {
	redisKey := oc.redisKey(groupKey, cacheKey)

	result, err := oc.redis.Get(context.Background(), redisKey).Bytes()
	if err == redis.Nil {
		return "", false, nil
	} else if err != nil {
		return "", false, err
	}

	return string(result), true, nil
}
