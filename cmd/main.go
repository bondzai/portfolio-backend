package main

import (
	"log/slog"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/handlers"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/bondzai/portfolio-backend/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var cfg = config.LoadConfig()

func initMongoDB() repositories.MongoDBClient {
	mongoClient, err := repositories.NewMongoDBClient(
		cfg.MongoUrl,
		cfg.MongoDB,
	)

	if err != nil {
		slog.Error("Failed to connect to MongoDB", err)
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

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     cfg.CorsOrigin,
		AllowHeaders:     cfg.CorsHeader,
		ExposeHeaders:    "Content-Length",
	}))

	mongoRepo := initMongoDB()

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
		slog.Error("Failed to start server", err)
	}
}
