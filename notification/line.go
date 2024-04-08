package notification

import "fmt"

type LineNotifyService struct {
	AccessToken string
}

func (l *LineNotifyService) SendNotification(payload map[string]interface{}) error {
	fmt.Println("Sending notification to Line Notify:", payload)
	return nil
}
