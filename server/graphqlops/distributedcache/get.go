package distributedcache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func Get(ctx context.Context, cacheKey string) (any, bool, error) {
	key := fmtKey(cacheKey)

	result, err := client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	var value any

	err = json.Unmarshal(result, &value)
	if err != nil {
		return value, false, err
	}

	return value, true, nil
}
