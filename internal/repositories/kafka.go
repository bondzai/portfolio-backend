package repositories

import "github.com/bondzai/portfolio-backend/pkg/kafka"

type (
	KafkaRepository interface {
		IsConnected() bool
		Publish(topic string, message interface{}) error
		Subscribe(topics []string, offsetOption int, consumerGroup string) (<-chan KafKaMessage, error)
		Close() error
	}

	KafKaMessage = *kafka.Message

	kafkaRepository struct {
		client kafka.Client
	}
)

func NewKafkaRepository(client kafka.Client) KafkaRepository {
	return &kafkaRepository{
		client: client,
	}
}

func (r *kafkaRepository) IsConnected() bool {
	return r.client.IsConnected()
}

func (r *kafkaRepository) Publish(topic string, message interface{}) error {
	return r.client.Publish(topic, message)
}

func (r *kafkaRepository) Subscribe(topics []string, offsetOption int, consumerGroup string) (<-chan KafKaMessage, error) {
	return r.client.Subscribe(topics, offsetOption, consumerGroup)
}

func (r *kafkaRepository) Close() error {
	return r.client.Close()
}
