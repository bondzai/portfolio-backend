package main

import (
	"log"
	"os"

	"github.com/bondzai/test/handlers"
	"github.com/bondzai/test/interfaces"
	"github.com/bondzai/test/userconnection"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/robfig/cron/v3"
)

func main() {
	app := fiber.New()

	configureCORS(app)

	mongoClient := initializeMongoDB()
	userManager := userconnection.NewUserConnectionManager(mongoClient)
	setupWebSocketRoutes(app, userManager)

	handlers.RegisterEndpoints(app)

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

func initializeMongoDB() interfaces.MongoDBClientInterface {
	mongoClient, err := interfaces.NewMongoDBClient(os.Getenv("GO_MONGODB_URL"), "portfolio", "usage")
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient
}

func setupWebSocketRoutes(app *fiber.App, ucm *userconnection.UserConnectionManager) {
	app.Get("/ws", websocket.New(ucm.HandleConnection))
}

func startCronJob(ucm *userconnection.UserConnectionManager) {
	c := cron.New()
	c.AddFunc("59 23 * * *", func() {
		ucm.ResetDailyUserCount()
	})
	c.Start()
	defer c.Stop()

	log.Println("cron started...")
}
