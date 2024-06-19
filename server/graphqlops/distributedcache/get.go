package distributedcache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func Get[T any](ctx context.Context, cacheKey string) (*T, bool, error) {
	key := fmtKey(cacheKey)

	result, err := client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	var value T
	err = json.Unmarshal(result, &value)
	if err != nil {
		return nil, false, err
	}

	return &value, true, nil
}
