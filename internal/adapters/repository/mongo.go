package repository

import (
	"context"
	"log"
	"time"

	"github.com/bondzai/portfolio-backend/internal/core/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClientInterface interface {
	SetDataToMongo(data *models.TotalUsers)
}

type MongoDBClient struct {
	client *mongo.Client
	db     *mongo.Database
	coll   *mongo.Collection
}

func NewMongoDBClient(connectionString, dbName, collectionName string) (*MongoDBClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	coll := db.Collection(collectionName)

	return &MongoDBClient{
		client: client,
		db:     db,
		coll:   coll,
	}, nil
}

func (mc *MongoDBClient) SetDataToMongo(data *models.TotalUsers) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := mc.coll.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
	}
}
