package repository

import (
	"context"
	"time"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient interface {
	InsertOne(collectionName string, data *domain.TotalUsers) error
	InsertMany(collectionName string, data []interface{}) error
	ReadCerts() ([]domain.Certification, error)
	ReadProjects() ([]domain.Project, error)
	ReadSkills() ([]domain.Skill, error)
}

type mongoDBClient struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDBClient(connectionString, dbName string) (*mongoDBClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &mongoDBClient{
		client: client,
		db:     db,
	}, nil
}

func (mc *mongoDBClient) InsertOne(collectionName string, data *domain.TotalUsers) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mc.db.Collection(collectionName)

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (mc *mongoDBClient) InsertMany(collectionName string, data []interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mc.db.Collection(collectionName)

	_, err := collection.InsertMany(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (mc *mongoDBClient) ReadCerts() ([]domain.Certification, error) {
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

	var certifications []domain.Certification
	for _, result := range results {
		var certification domain.Certification
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &certification)
		certifications = append(certifications, certification)
	}

	return certifications, nil
}

func (mc *mongoDBClient) ReadProjects() ([]domain.Project, error) {
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

	var projects []domain.Project
	for _, result := range results {
		var project domain.Project
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &project)
		projects = append(projects, project)
	}

	return projects, nil
}

func (mc *mongoDBClient) ReadSkills() ([]domain.Skill, error) {
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

	var skills []domain.Skill
	for _, result := range results {
		var skill domain.Skill
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &skill)
		skills = append(skills, skill)
	}

	return skills, nil
}
