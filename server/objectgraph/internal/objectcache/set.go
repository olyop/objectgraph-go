package objectcache

import (
	"context"
	"encoding/json"
	"time"
)

func (oc *ObjectCache) Set(
	groupKey string,
	cacheKey string,
	object map[string]any,
	ttl time.Duration,
) error {
	objectCache := oc.objectCache[groupKey]

	// lock the object so only one goroutine can set it at a time
	objectLocker := oc.objectLocker(groupKey, cacheKey)
	objectLocker.Lock()
	defer objectLocker.Unlock()

	// set the object in redis
	err := oc.redisSet(groupKey, cacheKey, object, ttl)
	if err != nil {
		objectCache.Delete(cacheKey)

		return err
	}

	// set the object in the cache
	oc.inmemorySet(groupKey, cacheKey, object, ttl)

	return nil
}

func (oc *ObjectCache) inmemorySet(
	groupKey string,
	cacheKey string,
	object any,
	ttl time.Duration,
) {
	objectCache := oc.objectCache[groupKey]

	objectCache.Set(cacheKey, object, ttl)
}

func (oc *ObjectCache) redisSet(
	groupKey string,
	cacheKey string,
	object any,
	ttl time.Duration,
) error {
	redisKey := oc.redisKey(groupKey, cacheKey)

	json, err := json.Marshal(object)
	if err != nil {
		return err
	}

	err = oc.redis.Set(context.Background(), redisKey, json, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}
