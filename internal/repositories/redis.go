package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
}

type redisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(addr, password string, db int) *redisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &redisClient{
		client: rdb,
		ctx:    context.Background(),
	}
}
