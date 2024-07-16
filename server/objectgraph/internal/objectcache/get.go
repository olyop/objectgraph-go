package objectcache

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/fatih/structs"
	"github.com/redis/go-redis/v9"
)

func (oc *ObjectCache) Get(
	groupKey string,
	reflectType reflect.Type,
	cacheKey string,
	ttl time.Duration,
) (map[string]any, bool, error) {
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
	object, exists, err := oc.redisGet(groupKey, reflectType, cacheKey)
	if err != nil {
		return nil, false, err
	}

	if exists {
		oc.inmemorySet(groupKey, cacheKey, object, ttl)

		return object, true, nil
	}

	return nil, false, nil
}

func (oc *ObjectCache) inmemoryGet(groupKey string, cacheKey string) (map[string]any, bool) {
	objectCache := oc.objectCache[groupKey]

	object, exists := objectCache.Get(cacheKey)
	if exists {
		return object.(map[string]any), true
	}

	return nil, false
}

func (oc *ObjectCache) redisGet(groupKey string, reflectType reflect.Type, cacheKey string) (map[string]any, bool, error) {
	redisKey := oc.redisKey(groupKey, cacheKey)
	data, err := oc.redis.Get(context.Background(), redisKey).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	result := reflect.New(reflectType).Interface()
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, false, err
	}

	object := structs.Map(result)

	return object, true, nil
}
