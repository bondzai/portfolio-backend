// handlers/handlers.go

package handlers

import (
	"fmt"

	"github.com/bondzai/test/data"
	"github.com/bondzai/test/notification"
	"github.com/bondzai/test/utils"
	"github.com/gofiber/fiber/v2"
)

func endpointHandler(sendNotification func([]string, map[string]interface{}), platforms []string, responseHandler func(*fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		payload := map[string]interface{}{
			"event_type": c.Path(),
			"data": map[string]interface{}{
				"ip": c.IP(),
			},
		}
		sendNotification(platforms, payload)
		return responseHandler(c)
	}
}

var WebhookURL = utils.GetEnv("DISCORD_WEBHOOK_URL", "")
var LineNotifyAccessToken = utils.GetEnv("LINE_NOTIFY_TOKEN", "")

func RegisterEndpoints(app *fiber.App) {
	notificationServices := map[string]notification.NotificationService{
		"discord": &notification.DiscordWebhookService{
			WebhookURL: WebhookURL,
		},
		"line": &notification.LineNotifyService{
			AccessToken: LineNotifyAccessToken,
		},
	}

	endpointPlatforms := map[string][]string{
		"/":               {""},
		"/skills":         {""},
		"/certifications": {""},
		"/projects":       {""},
		"/wakatime":       {""},
	}

	endpoints := map[string]func(*fiber.Ctx) error{
		"/": func(c *fiber.Ctx) error {
			return c.SendString("Ok")
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

	sendNotification := func(platforms []string, payload map[string]interface{}) {
		if len(platforms) == 0 {
			return
		}

		for _, platform := range platforms {
			if service, ok := notificationServices[platform]; ok {
				if err := service.SendNotification(payload); err != nil {
					fmt.Printf("Error sending notification to %s: %v\n", platform, err)
				}
			} else {
				fmt.Printf("Notification service for platform %s not found\n", platform)
			}
		}
	}

	for path, handler := range endpoints {
		platforms := endpointPlatforms[path]
		app.Get(path, endpointHandler(sendNotification, platforms, handler))
	}
}
