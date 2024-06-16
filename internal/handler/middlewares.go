package handler

import (
	"net/http"
	"strings"

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

func AuthTokenMiddleware(allowedToken string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token != allowedToken {
			return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
		}

		return c.Next()
	}
}
