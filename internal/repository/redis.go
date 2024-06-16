package repository

import (
	"context"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	InsertOne(collectionName string, data *domain.TotalUsers) error
	InsertMany(collectionName string, data []interface{}) error
	ReadCerts() ([]domain.Certification, error)
	ReadProjects() ([]domain.Project, error)
	ReadSkills() ([]domain.Skill, error)
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
