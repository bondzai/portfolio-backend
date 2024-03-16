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

	app.Get("/certifications", func(c *fiber.Ctx) error {
		return c.JSON(data.Certifications)
	})

	app.Get("/projects", func(c *fiber.Ctx) error {
		return c.JSON(data.Projects)
	})

	app.Listen(":10000")
}
