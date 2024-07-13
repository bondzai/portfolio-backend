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

	kafkaClient.Publish("myTopic", "test...")

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
		"127.0.0.1:6379",
		"",
		0,
	)

	return redisClient
}

func initKafka() kafka.Client {
	kafkaClient, err := kafka.NewClient(kafka.Config{
		Brokers: []string{"localhost:9092"},
		// For Cloud Karafka, use the following configuration
		// Brokers:  []string{
		// 	"broker1.cloudkarafka.com:9094",
		// 	"broker2.cloudkarafka.com:9094",
		// 	"broker3.cloudkarafka.com:9094",
		// },
		// Username: "your_sasl_username",
		// Password: "your_sasl_password",
		// UseTLS:   true,
	})
	if err != nil {
		log.Fatalf("Failed to setup Kafka client %v", err)
	}

	return kafkaClient
}
