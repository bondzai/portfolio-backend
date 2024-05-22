package main

import (
	"log"
	"os"

	repository "github.com/bondzai/portfolio-backend/internal/adapters/repository"
	usecases "github.com/bondzai/portfolio-backend/internal/core"
	"github.com/bondzai/portfolio-backend/internal/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"

	"github.com/robfig/cron/v3"
)

func main() {
	app := fiber.New()

	configureCORS(app)

	mongoClient := initMongoDB()
	userManager := usecases.NewManager(mongoClient)

	setupWebSocketRoutes(app, userManager)
	setupAPIRoutes(app)
	startCronJob(userManager)

	app.Listen(":" + os.Getenv("GO_PORT"))
}

func initMongoDB() repository.MongoDBClientInterface {
	mongoClient, err := repository.NewMongoDBClient(
		os.Getenv("GO_MONGODB_URL"),
		os.Getenv("GO_MONGODB_DB"),
		os.Getenv("GO_MONGODB_COL"),
	)

	if err != nil {
		log.Println(err)
	}

	return mongoClient
}

func configureCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("GO_CORS_ORIGINS"),
		AllowHeaders:     os.Getenv("GO_CORS_HEADERS"),
		ExposeHeaders:    "Content-Length",
		AllowCredentials: false,
	}))
}

func setupWebSocketRoutes(app *fiber.App, userManager *usecases.Manager) {
	app.Get("/ws", websocket.New(userManager.HandleConnection))
}

func setupAPIRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})
	app.Get("/skills", func(c *fiber.Ctx) error {
		return c.JSON(models.Skills)
	})
	app.Get("/certifications", func(c *fiber.Ctx) error {
		return c.JSON(models.Certifications)
	})
	app.Get("/projects", func(c *fiber.Ctx) error {
		return c.JSON(models.Projects)
	})
	app.Get("/wakatime", func(c *fiber.Ctx) error {
		return c.JSON(models.Wakatime)
	})
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
