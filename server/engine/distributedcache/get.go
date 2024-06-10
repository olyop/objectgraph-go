package distributedcache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func Get[R any](ctx context.Context, cacheKey string) (R, bool, error) {
	key := fmtKey(cacheKey)

	var value R

	result, err := client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return value, false, nil
	} else if err != nil {
		return value, false, err
	}

	err = json.Unmarshal(result, &value)
	if err != nil {
		return value, false, err
	}

	return value, true, nil
}
