package notification

import "fmt"

type DiscordWebhookService struct {
	WebhookURL string
}

func (d *DiscordWebhookService) SendNotification(payload map[string]interface{}) error {
	fmt.Println("Sending notification to Discord webhook:", payload)
	return nil
}
