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

func initializeMongoDB() interfaces.MongoDBClientInterface {
	mongoClient, err := interfaces.NewMongoDBClient(os.Getenv("GO_MONGODB_URL"), "portfolio", "usage")
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient
}

func main() {
	app := fiber.New()

	configureCORS(app)

	mongoClient := initializeMongoDB()
	ucm := userconnection.NewUserConnectionManager(mongoClient)

	setupWebSocketRoutes(app, ucm)

	handlers.RegisterEndpoints(app)

	startCronJob(ucm)

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
