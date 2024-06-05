package redis

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var Client *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	Client = client
}
