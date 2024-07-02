package objectcache

import (
	"context"
	"encoding/json"
	"time"
)

func (oc *ObjectCache) Set(
	groupKey string,
	cacheKey string,
	value any,
	ttl time.Duration,
) error {
	cacheGroup := oc.getCacheGroup(groupKey)

	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	json := string(jsonData)

	objectlocker := oc.getObjectLocker(groupKey, cacheKey)
	objectlocker.Lock()
	defer objectlocker.Unlock()

	// set the object in the cache
	cacheGroup.Set(cacheKey, json, ttl)

	// set the object in redis
	err = oc.redisSet(groupKey, cacheKey, json, ttl)
	if err != nil {
		cacheGroup.Delete(cacheKey)

		return err
	}

	return nil
}

func (oc *ObjectCache) redisSet(
	groupKey string,
	cacheKey string,
	value string,
	ttl time.Duration,
) error {
	redisKey := oc.redisKey(groupKey, cacheKey)

	_, err := oc.redis.Set(context.Background(), redisKey, value, ttl).Result()
	if err != nil {
		return err
	}

	return nil
}
