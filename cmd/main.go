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

	wsSerivce := services.NewWsService(mongoReo)

	handler := handler.NewHttpHandler(
		certService,
		projectService,
		skillService,
		wakaService,
	)

	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     cfg.CorsOrigin,
		AllowHeaders:     cfg.CorsHeader,
		ExposeHeaders:    "Content-Length",
	}))

	app.Get("/", handler.HealthCheck)
	app.Get("/certifications", handler.GetCerts)
	app.Get("/projects", handler.GetProjects)
	app.Get("/skills", handler.GetSkills)
	app.Get("/wakatime", handler.GetWakaStats)

	app.Get("/ws", websocket.New(wsSerivce.HandleConnection))

	wsSerivce.StartCronJob()

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
