package main

import (
	"github.com/bondzai/test/data"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/skills", func(c *fiber.Ctx) error {
		return c.JSON(data.Skills)
	})

	app.Listen(":10000")
}
