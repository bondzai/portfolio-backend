package main

import (
	"github.com/bondzai/test/data"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

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

	app.Listen(":10000")
}
