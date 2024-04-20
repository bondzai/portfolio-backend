package main

import (
	"os"

	"github.com/bondzai/test/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("GO_CORS_ORIGINS"),
		AllowHeaders:     os.Getenv("GO_CORS_HEADERS"),
		AllowCredentials: false,
	}))

	handlers.RegisterEndpoints(app)

	app.Listen(":" + os.Getenv("GO_PORT"))
}
