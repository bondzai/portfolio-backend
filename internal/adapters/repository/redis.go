package repository

import (
	"context"

	"github.com/bondzai/portfolio-backend/internal/core/models"
	"github.com/go-redis/redis/v8"
)

type RedisClientInterface interface {
	InsertOne(collectionName string, data *models.TotalUsers) error
	InsertMany(collectionName string, data []interface{}) error
	ReadCerts() ([]models.Certification, error)
	ReadProjects() ([]models.Project, error)
	ReadSkills() ([]models.Skill, error)
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
