package objectcache

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

type ObjectCache struct {
	redis  *redis.Client
	prefix string

	cacheGroups *cacheGroups
}

func New(prefix string, redisOpt *redis.Options) (*ObjectCache, error) {
	client := redis.NewClient(redisOpt)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	if prefix == "" {
		prefix = "objectcache"
	} else {
		prefix = prefix + ":objectcache"
	}

	oc := &ObjectCache{
		redis:  client,
		prefix: prefix,
		cacheGroups: &cacheGroups{
			lock:   &sync.RWMutex{},
			groups: make(map[string]*cacheGroup),
		},
	}

	return oc, nil
}

func (oc *ObjectCache) Close() {
	oc.redis.Close()
}
