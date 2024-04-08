package notification

type NotificationService interface {
	SendNotification(payload map[string]interface{}) error
}
