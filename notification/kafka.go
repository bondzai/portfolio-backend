package notification

import "fmt"

type KafkaService struct {
	BrokerURL string
}

func (k *KafkaService) SendNotification(payload map[string]interface{}) error {
	fmt.Println("Sending notification to Kafka:", payload)
	return nil
}
