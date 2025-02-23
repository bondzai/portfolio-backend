package infrastructures

import (
	"github.com/bondzai/portfolio-backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiber() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     config.AppConfig.CorsOrigin,
		AllowHeaders:     config.AppConfig.CorsHeader,
		ExposeHeaders:    "Content-Length",
	}))

	return app
}
