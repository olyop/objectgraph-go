package objectcache

import (
	"context"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

type ObjectCache struct {
	redis               *redis.Client
	prefix              string
	cleanupInterval     time.Duration
	cacheGroups         map[string]*cache.Cache
	cacheGroupsLock     *sync.Mutex
	objectLockers       map[string]map[string]*sync.Mutex
	objectLockersLock   *sync.Mutex
	objectLockersLocker map[string]*sync.Mutex
	clearLock           *sync.Mutex
}

func New(config *Configuration) (*ObjectCache, error) {
	client := redis.NewClient(config.Redis)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	oc := &ObjectCache{
		redis:               client,
		prefix:              config.Prefix + ":objectcache",
		cleanupInterval:     time.Minute,
		clearLock:           &sync.Mutex{},
		cacheGroups:         make(map[string]*cache.Cache),
		cacheGroupsLock:     &sync.Mutex{},
		objectLockers:       make(map[string]map[string]*sync.Mutex),
		objectLockersLock:   &sync.Mutex{},
		objectLockersLocker: make(map[string]*sync.Mutex),
	}

	return oc, nil
}

func (oc *ObjectCache) Close() {
	oc.redis.Close()
}

type Configuration struct {
	Prefix string
	Redis  *redis.Options
}
