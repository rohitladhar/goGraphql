package grpcstreaming

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context) *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password:"",
		DB: 0,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Errorf("failed to ping redis client: %w", err))
	}
	return client
}