package main

import (
	"log"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/adapters/handler"
	"github.com/bondzai/portfolio-backend/internal/adapters/repository"
	"github.com/bondzai/portfolio-backend/internal/core/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var cfg = config.LoadConfig()

func main() {
	app := fiber.New()

	mockRepo := repository.NewMock()
	mongoReo := initMongoDB()

	certService := services.NewCertService(mockRepo)
	projectService := services.NewProjectService(mockRepo)
	skillService := services.NewSkillService(mockRepo)
	wakaService := services.NewStatService()
	websocketService := services.NewWsService(mongoReo)

	restHandler := handler.NewHttpHandler(
		certService,
		projectService,
		skillService,
		wakaService,
	)

	websocketHandler := handler.NewWsHandler(websocketService)

	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     cfg.CorsOrigin,
		AllowHeaders:     cfg.CorsHeader,
		ExposeHeaders:    "Content-Length",
	}))

	app.Get("/", restHandler.HealthCheck)
	app.Get("/certifications", restHandler.GetCerts)
	app.Get("/projects", restHandler.GetProjects)
	app.Get("/skills", restHandler.GetSkills)
	app.Get("/wakatime", restHandler.GetWakaStats)

	app.Get("/ws", websocket.New(websocketHandler.HandleConnection))

	websocketService.StartCronJob()

	app.Listen(":" + cfg.Port)
}

func initMongoDB() repository.MongoDBClientInterface {
	mongoReo, err := repository.NewMongoDBClient(
		cfg.MongoUrl,
		cfg.MongoDB,
		cfg.MongoCol,
	)

	if err != nil {
		log.Println(err)
	}

	return mongoReo
}
