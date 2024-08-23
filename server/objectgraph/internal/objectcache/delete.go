package objectcache

import (
	"context"
)

func (oc *ObjectCache) Delete(
	groupKey string,
	objectKey string,
	valueKey string,
) error {
	// delete the object from redis
	redisKey := oc.redisKey(groupKey, objectKey, valueKey)
	err := oc.redis.Del(context.Background(), redisKey).Err()
	if err != nil {
		return err
	}

	// delete the object value in the inmemory cache
	oc.deleteValue(groupKey, objectKey, valueKey)

	return nil
}
