package distributedcache

import "context"

func Clear() error {
	_, err := client.FlushAll(context.Background()).Result()
	return err
}
