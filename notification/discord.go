package notification

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type DiscordWebhookService struct {
	WebhookURL string
}

func (d *DiscordWebhookService) SendNotification(payload map[string]interface{}) error {
	data := payload["data"].(map[string]interface{})
	ip := data["ip"].(string)

	jsonData := fmt.Sprintf(`{"content": "%s"}`, ip)

	req, err := http.NewRequest("POST", d.WebhookURL, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request: %v", err)
		return err
	}

	defer res.Body.Close()
	return nil
}
