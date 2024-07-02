package objectcache

import (
	"sync"

	"github.com/patrickmn/go-cache"
)

func (oc *ObjectCache) getCacheGroup(groupKey string) *cache.Cache {
	cacheGroup, initialized := oc.cacheGroups[groupKey]

	if !initialized {
		oc.cacheGroupsLock.Lock()
		defer oc.cacheGroupsLock.Unlock()

		// check again if the cacheGroup is initialized
		cacheGroup, initialized = oc.cacheGroups[groupKey]
		if !initialized {
			cacheGroup = cache.New(cache.DefaultExpiration, oc.cleanupInterval)
			oc.cacheGroups[groupKey] = cacheGroup
		}
	}

	return cacheGroup
}

func (oc *ObjectCache) getObjectLocker(groupKey string, cacheKey string) *sync.Mutex {
	objectLockerGroup := oc.getObjectLockerGroup(groupKey)

	objectLocker, initialized := objectLockerGroup[cacheKey]

	if !initialized {
		objectlockerMu := oc.getObjectLockersLocker(groupKey)

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

func (oc *ObjectCache) getObjectLockerGroup(groupKey string) map[string]*sync.Mutex {
	objectLockerGroup, initialized := oc.objectLockers[groupKey]

	if !initialized {
		oc.objectLockersLock.Lock()
		defer oc.objectLockersLock.Unlock()

		// check again if the objectLockerGroup is initialized
		objectLockerGroup, initialized = oc.objectLockers[groupKey]
		if !initialized {
			objectLockerGroup = make(map[string]*sync.Mutex)
			oc.objectLockers[groupKey] = objectLockerGroup
		}
	}

	return objectLockerGroup
}

func (oc *ObjectCache) getObjectLockersLocker(groupKey string) *sync.Mutex {
	objectLockersLocker, initialized := oc.objectLockersLocker[groupKey]

	if !initialized {
		oc.objectLockersLock.Lock()
		defer oc.objectLockersLock.Unlock()

		// check again if the objectLockersLocker is initialized
		objectLockersLocker, initialized = oc.objectLockersLocker[groupKey]
		if !initialized {
			objectLockersLocker = &sync.Mutex{}
			oc.objectLockersLocker[groupKey] = objectLockersLocker
		}
	}

	return objectLockersLocker

}

func (oc *ObjectCache) redisKey(groupKey string, cacheKey string) string {
	return oc.prefix + ":" + groupKey + ":" + cacheKey
}
