package pkg

import (
	"github.com/redis/go-redis/v9"
	"os"
)

func Redis(db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
	return client
}
