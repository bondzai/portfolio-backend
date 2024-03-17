package middlewares

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

// AuthTokenMiddleware is a middleware function to authenticate requests using tokens
func AuthTokenMiddleware(allowedToken string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header from the request
		authHeader := c.Get("Authorization")

		// Check if the Authorization header is missing or doesn't start with "Bearer"
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
		}

		// Extract the token from the Authorization header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Check if the token matches the allowed token
		if token != allowedToken {
			return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
		}

		// If the token matches, allow the request to proceed
		return c.Next()
	}
}
