package bootstrap

import (
	"context"
	"fmt"

	redis "github.com/redis/go-redis/v9"
)

func NewRedisClient(env *Env) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort),
		Password: env.RedisPass,
		DB:       env.RedisDB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return client
}
