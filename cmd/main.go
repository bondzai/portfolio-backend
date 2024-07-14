package main

import (
	"log"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/handlers"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/bondzai/portfolio-backend/internal/usecases"
	"github.com/bondzai/portfolio-backend/pkg/kafka"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var cfg = config.LoadConfig()

const kafkaDefaultTopic = "uzhfeczb-default"

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     cfg.CorsOrigin,
		AllowHeaders:     cfg.CorsHeader,
		ExposeHeaders:    "Content-Length",
	}))

	mongoRepo := initMongoDB()
	kafkaClient := initKafka()

	kafkaClient.Publish(kafkaDefaultTopic, "test...")

	certService := usecases.NewCertService(mongoRepo)
	projectService := usecases.NewProjectService(mongoRepo)
	skillService := usecases.NewSkillService(mongoRepo)
	wakaService := usecases.NewStatService()
	websocketService := usecases.NewWsService(mongoRepo)

	restHandler := handlers.NewHttpHandler(
		certService,
		projectService,
		skillService,
		wakaService,
	)

	websocketHandler := handlers.NewWsHandler(websocketService)

	app.Get("/", restHandler.HealthCheck)
	app.Get("/certifications", restHandler.GetCerts)
	app.Get("/projects", restHandler.GetProjects)
	app.Get("/skills", restHandler.GetSkills)
	app.Get("/wakatime", restHandler.GetWakaStats)

	app.Get("/ws", websocket.New(websocketHandler.HandleConnection))

	websocketService.StartCronJob()

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to set server %v", err)
	}
}

func initMongoDB() repositories.MongoDBClient {
	mongoClient, err := repositories.NewMongoDBClient(
		cfg.MongoUrl,
		cfg.MongoDB,
	)

	if err != nil {
		log.Fatalf("Failed to conect to MongoDB %v", err)
	}

	return mongoClient
}

func initRedis() repositories.RedisClient {
	redisClient := repositories.NewRedisClient(
		cfg.RedisUrl,
		cfg.RedisPass,
		cfg.RedisDb,
	)

	return redisClient
}

func initKafka() kafka.Client {
	kafkaClient, err := kafka.NewClient(kafka.Config{
		Brokers:          cfg.KafkaBroker,
		Username:         cfg.KafkaUser,
		Password:         cfg.KafkaPass,
		Mechanism:        "SCRAM-SHA-512",
		SecurityProtocol: "SASL_SSL",
	})
	if err != nil {
		log.Fatalf("Failed to setup Kafka client %v", err)
	}

	return kafkaClient
}
