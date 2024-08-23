package objectcache

import (
	"sync"
	"time"
)

type cacheGroups struct {
	lock   *sync.RWMutex
	groups map[string]*cacheGroup
}
type cacheGroup struct {
	lock    *sync.RWMutex
	objects map[string]*cacheObject
}
type cacheObject struct {
	lock   *sync.RWMutex
	values map[string]cacheValue
}
type cacheValue struct {
	lock    *sync.Mutex
	value   map[string]any
	expires time.Time
}

func (oc *ObjectCache) getValue(
	groupKey string,
	objectKey string,
	valueKey string,
) cacheValue {
	objectCache := oc.getObject(groupKey, objectKey)

	objectCache.lock.RLock()
	valueCache, exists := objectCache.values[valueKey]
	objectCache.lock.RUnlock()

	if !exists {
		objectCache.lock.Lock()
		valueCache = cacheValue{
			lock: &sync.Mutex{},
		}
		objectCache.values[valueKey] = valueCache
		objectCache.lock.Unlock()
	}

	if !valueCache.expires.IsZero() && time.Now().After(valueCache.expires) {
		valueCache.lock.Lock()
		valueCache.value = nil
		valueCache.expires = time.Time{}
		valueCache.lock.Unlock()
	}

	return valueCache
}

func (oc *ObjectCache) setValue(
	groupKey string,
	objectKey string,
	valueKey string,
	ttl time.Duration,
	value map[string]any,
) {
	objectCache := oc.getObject(groupKey, objectKey)

	expires := time.Now().Add(ttl)

	objectCache.lock.RLock()
	valueCache, exists := objectCache.values[valueKey]
	objectCache.lock.RUnlock()

	if exists {
		valueCache.lock.Lock()
		valueCache.value = value
		valueCache.expires = expires
		valueCache.lock.Unlock()
	} else {
		valueCache = cacheValue{
			lock:    &sync.Mutex{},
			value:   value,
			expires: expires,
		}

		objectCache.lock.Lock()
		objectCache.values[valueKey] = valueCache
		objectCache.lock.Unlock()
	}
}

func (oc *ObjectCache) deleteValue(
	groupKey string,
	objectKey string,
	valueKey string,
) {
	objectCache := oc.getObject(groupKey, objectKey)

	objectCache.lock.Lock()
	delete(objectCache.values, valueKey)
	objectCache.lock.Unlock()
}

func (oc *ObjectCache) getObject(
	groupKey string,
	objectKey string,
) *cacheObject {
	groupCache := oc.getGroup(groupKey)

	// read object
	groupCache.lock.RLock()
	objectCache, exists := groupCache.objects[objectKey]
	groupCache.lock.RUnlock()

	if !exists {
		// initialize object
		groupCache.lock.Lock()
		objectCache = &cacheObject{
			lock:   &sync.RWMutex{},
			values: make(map[string]cacheValue),
		}
		groupCache.objects[objectKey] = objectCache
		groupCache.lock.Unlock()
	}

	return objectCache
}

func (oc *ObjectCache) getGroup(groupKey string) *cacheGroup {
	// read group
	oc.cacheGroups.lock.RLock()
	groupCache, exists := oc.cacheGroups.groups[groupKey]
	oc.cacheGroups.lock.RUnlock()

	if !exists {
		// initialize group
		oc.cacheGroups.lock.Lock()
		groupCache = &cacheGroup{
			lock:    &sync.RWMutex{},
			objects: make(map[string]*cacheObject),
		}
		oc.cacheGroups.groups[groupKey] = groupCache
		oc.cacheGroups.lock.Unlock()
	}

	return groupCache
}
