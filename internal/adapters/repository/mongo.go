package repository

import (
	"context"
	"time"

	"github.com/bondzai/portfolio-backend/internal/core/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClientInterface interface {
	InsertTotalUsers(collectionName string, data *models.TotalUsers) error
	InsertMany(collectionName string, data []interface{}) error
}

type MongoDBClient struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDBClient(connectionString, dbName string) (*MongoDBClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &MongoDBClient{
		client: client,
		db:     db,
	}, nil
}

func (mc *MongoDBClient) InsertTotalUsers(collectionName string, data *models.TotalUsers) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mc.db.Collection(collectionName)

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (mc *MongoDBClient) InsertMany(collectionName string, data []interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mc.db.Collection(collectionName)

	_, err := collection.InsertMany(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
