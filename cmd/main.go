package main

import (
	"flag"
	"log/slog"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/handler"
	"github.com/bondzai/portfolio-backend/internal/repository"
	"github.com/bondzai/portfolio-backend/internal/usecase"
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

	var err error

	certifications, _ := mockRepo.ReadCerts()
	err = mongoRepo.InsertMany("certifications", utils.ConvertToInterfaceSlice(certifications))
	if err != nil {
		slog.Error("Error seeded certifications data", err)
	} else {
		slog.Info("Successfully seeded certifications data to MongoDB")
	}

	projects, _ := mockRepo.ReadProjects()
	err = mongoRepo.InsertMany("projects", utils.ConvertToInterfaceSlice(projects))
	if err != nil {
		slog.Error("Error seeded projects data", err)
	} else {
		slog.Info("Successfully seeded projects data to MongoDB")
	}

	skills, _ := mockRepo.ReadSkills()
	err = mongoRepo.InsertMany("skills", utils.ConvertToInterfaceSlice(skills))
	if err != nil {
		slog.Error("Error seeded skills data", err)
	} else {
		slog.Info("Successfully seeded skills data to MongoDB")
	}
}
