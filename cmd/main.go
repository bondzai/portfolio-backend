package main

import (
	"log"
	"os"

	"github.com/bondzai/test/data"
	"github.com/bondzai/test/interfaces"
	"github.com/bondzai/test/userconnection"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"

	"github.com/robfig/cron/v3"
)

func initMongoDB() interfaces.MongoDBClientInterface {
	mongoClient, err := interfaces.NewMongoDBClient(
		os.Getenv("GO_MONGODB_URL"),
		os.Getenv("GO_MONGODB_DB"),
		os.Getenv("GO_MONGODB_COL"),
	)

	if err != nil {
		log.Println(err)
	}

	return mongoClient
}

func main() {
	app := fiber.New()

	configureCORS(app)

	mongoClient := initMongoDB()
	userManager := userconnection.NewManager(mongoClient)

	setupWebSocketRoutes(app, userManager)
	setupAPIRoutes(app)
	startCronJob(userManager)

	app.Listen(":" + os.Getenv("GO_PORT"))
}

func configureCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("GO_CORS_ORIGINS"),
		AllowHeaders:     os.Getenv("GO_CORS_HEADERS"),
		ExposeHeaders:    "Content-Length",
		AllowCredentials: false,
	}))
}

func setupWebSocketRoutes(app *fiber.App, userManager *userconnection.Manager) {
	app.Get("/ws", websocket.New(userManager.HandleConnection))
}

func setupAPIRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})
	app.Get("/skills", func(c *fiber.Ctx) error {
		return c.JSON(data.Skills)
	})
	app.Get("/certifications", func(c *fiber.Ctx) error {
		return c.JSON(data.Certifications)
	})
	app.Get("/projects", func(c *fiber.Ctx) error {
		return c.JSON(data.Projects)
	})
	app.Get("/wakatime", func(c *fiber.Ctx) error {
		return c.JSON(data.Wakatime)
	})
}

func startCronJob(userManager *userconnection.Manager) {
	c := cron.New()
	c.AddFunc("59 23 * * *", func() {
		userManager.ResetDailyUserCount()
	})
	c.Start()
	defer c.Stop()

	log.Println("cron started...")
}
