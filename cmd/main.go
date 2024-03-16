package main

import (
	"net/http"

	"github.com/bondzai/test/data"
	"github.com/gofiber/fiber/v2"
)

func APIKeyAuthMiddleware(apiKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientAPIKey := c.Get("X-API-Key")

		if clientAPIKey != apiKey {
			return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
		}

		return c.Next()
	}
}

func main() {
	apiKey := "your-api-key"

	app := fiber.New()

	app.Use(APIKeyAuthMiddleware(apiKey))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
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

	app.Listen(":10000")
}
