package objectcache

import (
	"context"
	"sync"

	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

type ObjectCache struct {
	redis               *redis.Client
	prefix              string
	objectCache         objectCache
	objectLockers       objectLockers
	objectLockersLocker objectLockersLocker
}

type objectCache map[string]*cache.Cache
type objectLockers map[string]map[string]*sync.Mutex
type objectLockersLocker map[string]*sync.Mutex

func New(prefix string, r *redis.Options, typeNames []string) (*ObjectCache, error) {
	client := redis.NewClient(r)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	oc := &ObjectCache{
		redis:               client,
		prefix:              initPrefix(prefix),
		objectCache:         initObjectCache(typeNames),
		objectLockers:       initObjectLockers(typeNames),
		objectLockersLocker: initObjectLockersLocker(typeNames),
	}

	return oc, nil
}

func (oc *ObjectCache) Close() {
	oc.redis.Close()
}