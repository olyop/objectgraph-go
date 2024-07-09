package objectcache

import (
	"context"
)

func (oc *ObjectCache) Delete(
	groupKey string,
	cacheKey string,
) error {
	inmemoryCache := oc.objectCache[groupKey]

	objectlocker := oc.getObjectLocker(groupKey, cacheKey)
	objectlocker.Lock()
	defer objectlocker.Unlock()

	err := oc.redis.Del(context.Background(), oc.redisKey(groupKey, cacheKey)).Err()
	if err != nil {
		return err
	}

	inmemoryCache.Delete(cacheKey)

	return nil
}
