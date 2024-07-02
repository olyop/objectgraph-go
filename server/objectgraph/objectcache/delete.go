package objectcache

import (
	"context"
)

func (oc *ObjectCache) Delete(
	groupKey string,
	cacheKey string,
) error {
	cacheGroup := oc.getCacheGroup(groupKey)

	objectlocker := oc.getObjectLocker(groupKey, cacheKey)
	objectlocker.Lock()
	defer objectlocker.Unlock()

	err := oc.redis.Del(context.Background(), oc.redisKey(groupKey, cacheKey)).Err()
	if err != nil {
		return err
	}

	cacheGroup.Delete(cacheKey)

	return nil
}
