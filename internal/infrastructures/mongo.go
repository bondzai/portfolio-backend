package infrastructures

import (
	"log"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/repositories"
)

func NewMongo() repositories.MongoDBClient {
	mongoClient, err := repositories.NewMongoDBClient(
		config.AppConfig.MongoUrl,
		config.AppConfig.MongoDB,
	)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB %v", err)
	}

	return mongoClient
}
