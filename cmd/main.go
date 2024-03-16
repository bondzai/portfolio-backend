package main

import (
	"net/http"

	"github.com/bondzai/test/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// URLFilterMiddleware is a middleware function to filter requests based on requested URL
func URLFilterMiddleware(allowedURL string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestedURL := c.OriginalURL()

		// Check if the requested URL matches the allowed URL
		if requestedURL != allowedURL {
			return c.Status(http.StatusForbidden).SendString("Forbidden")
		}

		// Allow the request to continue if the requested URL matches the allowed URL
		return c.Next()
	}
}

func main() {
	allowedURL := "https://thejb.onrender.com"

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedURL,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(URLFilterMiddleware(allowedURL))

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
