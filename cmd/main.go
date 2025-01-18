package main

import (
	"log"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/handlers"
	"github.com/bondzai/portfolio-backend/internal/infrastructures"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/bondzai/portfolio-backend/internal/usecases"
	"github.com/bondzai/portfolio-backend/pkg/kafka"

	"github.com/gofiber/websocket/v2"
)

func main() {
	kafkaClient := initKafka()
	mongoClient := initMongoDB()

	kafkaRepository := repositories.NewKafkaRepository(kafkaClient)

	certService := usecases.NewCertService(mongoClient)
	projectService := usecases.NewProjectService(mongoClient)
	skillService := usecases.NewSkillService(mongoClient)
	wakaService := usecases.NewStatService()
	websocketService := usecases.NewWsService(mongoClient, kafkaRepository)
	websocketService.StartCronJob()

	restHandler := handlers.NewHttpHandler(
		certService,
		projectService,
		skillService,
		wakaService,
	)

	websocketHandler := handlers.NewWsHandler(websocketService)

	app := infrastructures.NewFiber()

	app.Get("/", restHandler.HealthCheck)
	app.Get("/certifications", restHandler.GetCerts)
	app.Get("/projects", restHandler.GetProjects)
	app.Get("/skills", restHandler.GetSkills)
	app.Get("/wakatime", restHandler.GetWakaStats)

	app.Get("/ws", websocket.New(websocketHandler.HandleConnection))

	if err := app.Listen(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to set server %v", err)
	}
}

func initMongoDB() repositories.MongoDBClient {
	mongoClient, err := repositories.NewMongoDBClient(
		config.AppConfig.MongoUrl,
		config.AppConfig.MongoDB,
	)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB %v", err)
	}

	return mongoClient
}

func initRedis() repositories.RedisClient {
	redisClient := repositories.NewRedisClient(
		config.AppConfig.RedisUrl,
		config.AppConfig.RedisPass,
		config.AppConfig.RedisDb,
	)

	return redisClient
}

func initKafka() kafka.Client {
	kafkaClient, err := kafka.NewClient(kafka.Config{
		Brokers:          config.AppConfig.KafkaBroker,
		Username:         config.AppConfig.KafkaUser,
		Password:         config.AppConfig.KafkaPass,
		Mechanism:        "SCRAM-SHA-512",
		SecurityProtocol: "SASL_SSL",
	})
	if err != nil {
		log.Fatalf("Failed to setup Kafka client %v", err)
	}

	return kafkaClient
}
