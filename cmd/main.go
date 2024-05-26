package main

import (
	"log"

	"github.com/bondzai/portfolio-backend/internal/handlers"
	"github.com/bondzai/portfolio-backend/internal/middlewares"
	"github.com/bondzai/portfolio-backend/internal/services/mongodb"
	"github.com/bondzai/portfolio-backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("initial started")

	mongodb.CheckConnection()
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     utils.GetEnv("GO_CORS_ORIGINS", ""),
		AllowHeaders:     utils.GetEnv("GO_CORS_HEADERS", "*"),
		AllowMethods:     utils.GetEnv("GO_CORS_METHODS", "*"),
		AllowCredentials: true,
	}))
	app.Use(middlewares.CustomAuth)

	dataHandler := handlers.NewDataHandler()

	app.Get("/:dataType", dataHandler.HandleData)

	app.Post("/flush-cache", middlewares.CustomExtraAuth, dataHandler.FlushCache)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("JB backend is running...")
	})

	port := utils.GetEnv("PORT", ":10000")
	log.Printf("Server is running on port %s", port)
	app.Listen(port)
}
