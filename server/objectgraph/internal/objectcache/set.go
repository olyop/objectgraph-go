package objectcache

import (
	"context"
	"encoding/json"
	"time"
)

func (oc *ObjectCache) Set(
	groupKey string,
	cacheKey string,
	object any,
	ttl time.Duration,
) error {
	inmemoryCache := oc.objectCache[groupKey]

	objectlocker := oc.getObjectLocker(groupKey, cacheKey)
	objectlocker.Lock()
	defer objectlocker.Unlock()

	// set the object in redis
	err := oc.redisSet(groupKey, cacheKey, object, ttl)
	if err != nil {
		inmemoryCache.Delete(cacheKey)

		return err
	}

	// set the object in the cache
	inmemoryCache.Set(cacheKey, object, ttl)

	return nil
}

func (oc *ObjectCache) redisSet(
	groupKey string,
	cacheKey string,
	object any,
	ttl time.Duration,
) error {
	redisKey := oc.redisKey(groupKey, cacheKey)

	json, err := redisSerializeObject(object)
	if err != nil {
		return err
	}

	err = oc.redis.Set(context.Background(), redisKey, json, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func redisSerializeObject(value any) (string, error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
