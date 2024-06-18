package distributedcache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var prefix string

func Connect(pre string, options *redis.Options) error {
	prefix = pre

	client = redis.NewClient(options)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	err := client.Close()
	if err != nil {
		panic(err)
	}
}
