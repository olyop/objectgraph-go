package objectcache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

func initializeCacheGroups(groupKeys []string) map[string]*cache.Cache {
	cacheGroups := make(map[string]*cache.Cache)

	for _, key := range groupKeys {
		cacheGroups[key] = cache.New(time.Second*60, time.Second*60)
	}

	return cacheGroups
}

func initializeObjectLockers(groupKeys []string) map[string]map[string]*sync.Mutex {
	cacheItemLockers := make(map[string]map[string]*sync.Mutex)

	for _, key := range groupKeys {
		cacheItemLockers[key] = make(map[string]*sync.Mutex)
	}

	return cacheItemLockers
}

func initializeObjectLockersMu(groupKeys []string) map[string]*sync.Mutex {
	cacheItemLockersMu := make(map[string]*sync.Mutex)

	for _, key := range groupKeys {
		cacheItemLockersMu[key] = &sync.Mutex{}
	}

	return cacheItemLockersMu
}
