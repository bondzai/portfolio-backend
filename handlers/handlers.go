// handlers/handlers.go

package handlers

import (
	"fmt"
	"time"

	"github.com/bondzai/test/data"
	"github.com/bondzai/test/notification"
	"github.com/gofiber/fiber/v2"
)

// Define a handler function for all endpoints
func endpointHandler(sendNotification func([]string, map[string]interface{}), platforms []string, responseHandler func(*fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		payload := map[string]interface{}{
			"event_type": c.Path(),
			"data": map[string]interface{}{
				"user_id":   "user123",
				"timestamp": time.Now(),
			},
		}
		sendNotification(platforms, payload)
		return responseHandler(c)
	}
}

// RegisterEndpoints registers the endpoints with the Fiber app
func RegisterEndpoints(app *fiber.App) {
	// Initialize notification services
	notificationServices := map[string]notification.NotificationService{
		"discord": &notification.DiscordWebhookService{
			WebhookURL: "https://discord.com/api/webhooks/your-discord-webhook-url",
		},
		"line": &notification.LineNotifyService{
			AccessToken: "your-line-access-token",
		},
	}

	// Map of endpoint paths and their corresponding notification platforms
	endpointPlatforms := map[string][]string{
		"/":               {"line"},
		"/skills":         {"discord"},
		"/certifications": {""}, // No notification
		"/projects":       {""}, // No notification
		"/wakatime":       {""}, // No notification
	}

	// Map of endpoint paths and their corresponding response handlers
	endpoints := map[string]func(*fiber.Ctx) error{
		"/": func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		},
		"/skills": func(c *fiber.Ctx) error {
			return c.JSON(data.Skills)
		},
		"/certifications": func(c *fiber.Ctx) error {
			return c.JSON(data.Certifications)
		},
		"/projects": func(c *fiber.Ctx) error {
			return c.JSON(data.Projects)
		},
		"/wakatime": func(c *fiber.Ctx) error {
			return c.JSON(data.Wakatime)
		},
	}

	// Define a handler function for sending notifications
	sendNotification := func(platforms []string, payload map[string]interface{}) {
		if len(platforms) == 0 {
			return // No platforms specified for notification
		}
		for _, platform := range platforms {
			if service, ok := notificationServices[platform]; ok {
				if err := service.SendNotification(payload); err != nil {
					fmt.Printf("Error sending notification to %s: %v\n", platform, err)
					// Handle error if necessary
				}
			} else {
				fmt.Printf("Notification service for platform %s not found\n", platform)
				// Handle error if necessary
			}
		}
	}

	// Register endpoints with different notification platforms and response handlers
	for path, handler := range endpoints {
		platforms := endpointPlatforms[path]
		app.Get(path, endpointHandler(sendNotification, platforms, handler))
	}
}
