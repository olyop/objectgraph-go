package objectcache

import (
	"sync"
)

func (oc *ObjectCache) objectLocker(groupKey string, cacheKey string) *sync.Mutex {
	objectLockerGroup := oc.objectLockers[groupKey]

	objectLocker, initialized := objectLockerGroup[cacheKey]

	if !initialized {
		objectlockerMu := oc.objectLockersLocker[groupKey]

		objectlockerMu.Lock()
		defer objectlockerMu.Unlock()

		// check again if the objectlocker is initialized
		objectLocker, initialized = objectLockerGroup[cacheKey]
		if !initialized {
			objectLocker = &sync.Mutex{}
			objectLockerGroup[cacheKey] = objectLocker
		}
	}

	return objectLocker
}

func (oc *ObjectCache) redisKey(groupKey string, cacheKey string) string {
	return oc.prefix + ":" + groupKey + ":" + cacheKey
}
