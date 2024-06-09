package engine

import (
	"context"
	"sync"
)

func getRequestMutext(ctx context.Context, cacheKey string) *sync.Mutex {
	m := ctx.Value(ResolverRequestMutexContextKey{}).(*sync.Map)

	mu, found := m.Load(cacheKey)
	if !found {
		mu = &sync.Mutex{}
		m.Store(cacheKey, mu)
	}

	return mu.(*sync.Mutex)
}

func getResolverMutex(ctx context.Context, cacheKey string) *sync.Mutex {
	m := ctx.Value(ResolverRetrieveMutexContextKey{}).(*sync.Map)

	mu, found := m.Load(cacheKey)
	if !found {
		mu = &sync.Mutex{}
		m.Store(cacheKey, mu)
	}

	return mu.(*sync.Mutex)
}
