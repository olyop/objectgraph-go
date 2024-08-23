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
	objectKey string,
	valueKey string,
	ttl time.Duration,
	reflectType reflect.Type,
) (map[string]any, bool, error) {
	// check the cache first
	valueCache := oc.getValue(groupKey, objectKey, valueKey)
	if valueCache.value != nil {
		return valueCache.value, true, nil
	}

	// not in the cache, lock the object value as only one goroutine should be fetching it
	valueCache.lock.Lock()
	defer valueCache.lock.Unlock()

	// check again for where another goroutine has already fetched the object
	valueCache = oc.getValue(groupKey, objectKey, valueKey)
	if valueCache.value != nil {
		return valueCache.value, true, nil
	}

	// fetch the object from redis
	value, exists, err := oc.redisGet(groupKey, objectKey, valueKey, reflectType)
	if err != nil {
		// error fetching from redis so delete the value from the cache
		oc.deleteValue(groupKey, objectKey, valueKey)
		return nil, false, err
	}
	if exists {
		// set the value in the inmemory cache
		oc.setValue(groupKey, objectKey, valueKey, ttl, value)
		return value, true, nil
	}

	// object not found
	return nil, false, nil
}

func (oc *ObjectCache) redisGet(
	groupKey string,
	objectKey string,
	valueKey string,
	reflectType reflect.Type,
) (map[string]any, bool, error) {
	redisKey := oc.redisKey(groupKey, objectKey, valueKey)

	data, err := oc.redis.Get(context.Background(), redisKey).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	// cast the object to the correct type
	result := reflect.New(reflectType).Interface()
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, false, err
	}

	return structs.Map(result), true, nil
}
