package objectcache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

func initPrefix(prefix string) string {
	if prefix == "" {
		return "objectcache"
	}

	return prefix + ":objectcache"
}

func initObjectCache(typeNames []string) objectCache {
	oc := make(objectCache)

	for _, typeName := range typeNames {
		oc[typeName] = cache.New(cache.NoExpiration, time.Minute)
	}

	return oc
}

func initObjectLockers(typeNames []string) objectLockers {
	ol := make(objectLockers)

	for _, typeName := range typeNames {
		ol[typeName] = make(map[string]*sync.Mutex)
	}

	return ol
}

func initObjectLockersLocker(typeNames []string) objectLockersLocker {
	oll := make(objectLockersLocker)

	for _, typeName := range typeNames {
		oll[typeName] = &sync.Mutex{}
	}

	return oll
}
