package handlers

import (
	"github.com/bondzai/portfolio-backend/internal/utils/errs"
	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errs.AppError:
		return c.Status(e.Code).JSON(fiber.Map{"error": e.Message})

	case error:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": e.Error()})

	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unexpected error"})
	}
}
