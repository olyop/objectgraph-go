package inmemorycache

import (
	"sync"

	"github.com/patrickmn/go-cache"
)

var state CacheState = CacheState{
	mu:     &sync.Mutex{},
	groups: make(CacheGroups),
}

type CacheState struct {
	mu     *sync.Mutex
	groups CacheGroups
}

type CacheGroups map[string]*cache.Cache
