package inmemorycache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func handleGroup(groupKey string) *cache.Cache {
	groupCache, initialized := state.groups[groupKey]

	if !initialized {
		state.mu.Lock()

		// Check again in case another goroutine initialized the group
		groupCache, initialized = state.groups[groupKey]
		if initialized {
			state.mu.Unlock()
			return groupCache
		}

		groupCache = cache.New(cache.NoExpiration, time.Minute)

		state.groups[groupKey] = groupCache

		state.mu.Unlock()
	}

	return groupCache
}
