package objectcache

import (
	"context"
	"encoding/json"
	"time"
)

func (oc *ObjectCache) Set(
	groupKey string,
	objectKey string,
	valueKey string,
	value map[string]any,
	ttl time.Duration,
) error {
	// set the object in redis
	err := oc.redisSet(groupKey, objectKey, valueKey, ttl, value)
	if err != nil {
		oc.deleteValue(groupKey, objectKey, valueKey)
		return err
	}

	// set the object value in the inmemory cache
	oc.setValue(groupKey, objectKey, valueKey, ttl, value)

	return nil
}

func (oc *ObjectCache) redisSet(
	groupKey string,
	objectKey string,
	valueKey string,
	ttl time.Duration,
	object map[string]any,
) error {
	redisKey := oc.redisKey(groupKey, objectKey, valueKey)

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
