package main

import (
	"log"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/adapters/handler"
	repository "github.com/bondzai/portfolio-backend/internal/adapters/repository"
	usecases "github.com/bondzai/portfolio-backend/internal/core"
	"github.com/bondzai/portfolio-backend/internal/core/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"

	"github.com/robfig/cron/v3"
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

	mockRepo := repository.NewMock()
	mongoReo := initMongoDB()

	certService := services.NewCertService(mockRepo)
	projectService := services.NewProjectService(mockRepo)
	skillService := services.NewSkillService(mockRepo)
	wakaService := services.NewStatService()

	userManager := usecases.NewManager(mongoReo)

	handler := handler.NewHttpHandler(
		certService,
		projectService,
		skillService,
		wakaService,
	)

	app.Get("/", handler.GetCerts)
	app.Get("/certifications", handler.GetCerts)
	app.Get("/projects", handler.GetProjects)
	app.Get("/skills", handler.GetSkills)
	app.Get("/wakatime", handler.GetWakaStats)

	setupWebSocketRoutes(app, userManager)

	startCronJob(userManager)

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

func setupWebSocketRoutes(app *fiber.App, userManager *usecases.Manager) {
	app.Get("/ws", websocket.New(userManager.HandleConnection))
}

func startCronJob(userManager *usecases.Manager) {
	c := cron.New()

	c.AddFunc("59 23 * * *", func() {
		userManager.ResetDailyUserCount()
	})

	c.Start()
	defer c.Stop()

	log.Println("cron started...")
}
