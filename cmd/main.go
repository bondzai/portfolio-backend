package main

import (
	"flag"
	"log/slog"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/handler"
	"github.com/bondzai/portfolio-backend/internal/repository"
	"github.com/bondzai/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var cfg = config.LoadConfig()

func initMongoDB() repository.MongoDBClientInterface {
	mongoRepo, err := repository.NewMongoDBClient(
		cfg.MongoUrl,
		cfg.MongoDB,
	)

	if err != nil {
		slog.Error("Failed to connect to MongoDB", err)
	}

	return mongoRepo
}

func main() {
	seedFlag := flag.Bool("seed", false, "Data seeding.")
	flag.Parse()

	if *seedFlag {
		runSeed()
	} else {
		runServer()
	}
}

func runServer() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     cfg.CorsOrigin,
		AllowHeaders:     cfg.CorsHeader,
		ExposeHeaders:    "Content-Length",
	}))

	mongoRepo := initMongoDB()

	certService := usecase.NewCertService(mongoRepo)
	projectService := usecase.NewProjectService(mongoRepo)
	skillService := usecase.NewSkillService(mongoRepo)
	wakaService := usecase.NewStatService()
	websocketService := usecase.NewWsService(mongoRepo)

	restHandler := handler.NewHttpHandler(
		certService,
		projectService,
		skillService,
		wakaService,
	)

	websocketHandler := handler.NewWsHandler(websocketService)

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
