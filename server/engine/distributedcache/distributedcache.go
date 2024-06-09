package distributedcache

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var prefix string

func Connect() {
	addr := os.Getenv("REDIS_URL")
	password := os.Getenv("REDIS_PASSWORD")
	prefix = os.Getenv("REDIS_PREFIX")

	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func Close() {
	client.Close()
}
