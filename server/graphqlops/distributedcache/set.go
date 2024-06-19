package distributedcache

import (
	"context"
	"encoding/json"
	"time"
)

func Set[T any](ctx context.Context, cacheKey string, value T, ttl time.Duration) error {
	key := fmtKey(cacheKey)

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, string(json), ttl).Err()
}
