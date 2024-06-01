package main

import (
	"flag"
	"log/slog"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/adapters/handler"
	"github.com/bondzai/portfolio-backend/internal/adapters/repository"
	"github.com/bondzai/portfolio-backend/internal/core/services"
	"github.com/bondzai/portfolio-backend/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var cfg = config.LoadConfig()

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

	certService := services.NewCertService(mongoRepo)
	projectService := services.NewProjectService(mongoRepo)
	skillService := services.NewSkillService(mongoRepo)
	wakaService := services.NewStatService()
	websocketService := services.NewWsService(mongoRepo)

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

func runSeed() {
	mockRepo := repository.NewMock()
	mongoRepo := initMongoDB()

	certifications, _ := mockRepo.ReadCerts()
	mongoRepo.InsertMany("certifications", utils.ConvertToInterfaceSlice(certifications))
	slog.Info("Successfully seeded certifications data to MongoDB")

	projects, _ := mockRepo.ReadProjects()
	mongoRepo.InsertMany("projects", utils.ConvertToInterfaceSlice(projects))
	slog.Info("Successfully seeded projects data to MongoDB")

	skills, _ := mockRepo.ReadSkills()
	mongoRepo.InsertMany("skills", utils.ConvertToInterfaceSlice(skills))
	slog.Info("Successfully seeded skills data to MongoDB")
}
