package repository

import (
	"context"
	"time"

	"github.com/bondzai/portfolio-backend/internal/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClientInterface interface {
	InsertOne(collectionName string, data *models.TotalUsers) error
	InsertMany(collectionName string, data []interface{}) error
	ReadCerts() ([]models.Certification, error)
	ReadProjects() ([]models.Project, error)
	ReadSkills() ([]models.Skill, error)
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

func (mc *MongoDBClient) InsertOne(collectionName string, data *models.TotalUsers) error {
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

func (mc *MongoDBClient) ReadCerts() ([]models.Certification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mc.db.Collection("certifications")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	var certifications []models.Certification
	for _, result := range results {
		var certification models.Certification
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &certification)
		certifications = append(certifications, certification)
	}

	return certifications, nil
}

func (mc *MongoDBClient) ReadProjects() ([]models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mc.db.Collection("projects")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	var projects []models.Project
	for _, result := range results {
		var project models.Project
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &project)
		projects = append(projects, project)
	}

	return projects, nil
}

func (mc *MongoDBClient) ReadSkills() ([]models.Skill, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mc.db.Collection("skills")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	var skills []models.Skill
	for _, result := range results {
		var skill models.Skill
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &skill)
		skills = append(skills, skill)
	}

	return skills, nil
}
