package main

import (
	"log"
	"os"

	"github.com/bondzai/test/handlers"
	"github.com/bondzai/test/interfaces"
	"github.com/bondzai/test/websocketmanager"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/robfig/cron/v3"
)

func main() {
	app := fiber.New()

	configureCORS(app)

	mongoClient := initializeMongoDB()
	wsm := websocketmanager.NewWebSocketManager(mongoClient)
	setupWebSocketRoutes(app, wsm)

	handlers.RegisterEndpoints(app)

	startCronJob(wsm)

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

func setupWebSocketRoutes(app *fiber.App, wsm *websocketmanager.WebSocketManager) {
	app.Get("/ws", websocket.New(wsm.HandleConnection))
}

func startCronJob(wsm *websocketmanager.WebSocketManager) {
	c := cron.New()
	c.AddFunc("59 23 * * *", func() {
		wsm.ResetDailyUserCount()
	})
	c.Start()
	defer c.Stop()

	log.Println("cron started...")
}
