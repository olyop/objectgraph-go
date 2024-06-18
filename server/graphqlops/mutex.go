package graphqlops

import (
	"context"
	"sync"
)

func loadResolverLocker(ctx context.Context, key any, cacheKey string) *sync.Mutex {
	m := ctx.Value(key).(*sync.Map)

	mu, found := m.Load(cacheKey)
	if !found {
		mu = &sync.Mutex{}
		m.Store(cacheKey, mu)
	}

	return mu.(*sync.Mutex)
}
