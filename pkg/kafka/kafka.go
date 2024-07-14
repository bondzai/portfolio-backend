package kafka

import (
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	OffsetFromLatest   = iota // Start from the latest offset processed
	OffsetFromEarliest        // Start from the beginning (oldest) offset
	Timeout            = 1
)

type (
	Client interface {
		IsConnected() bool
		Publish(topic string, message interface{}) error
		Subscribe(topics []string, offsetOption int, consumerGroup string) (<-chan *Message, error)
		Close() error
	}

	client struct {
		producer *kafka.Producer
		consumer *kafka.Consumer
		brokers  string
	}

	Message struct {
		Topic     string
		Partition int32
		Offset    int64
		Key       []byte
		Value     []byte
		Timestamp time.Time
	}

	Config struct {
		Brokers          string
		Username         string
		Password         string
		Mechanism        string
		SecurityProtocol string
	}
)

func NewClient(config Config) (Client, error) {
	producer, err := newProducer(config)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Kafka successfully.")

	return &client{
		producer: producer,
		brokers:  config.Brokers,
	}, nil
}

func newProducer(config Config) (*kafka.Producer, error) {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.Brokers,
	}

	if config.Username != "" && config.Password != "" {
		kafkaConfig.SetKey("sasl.username", config.Username)
		kafkaConfig.SetKey("sasl.password", config.Password)
		kafkaConfig.SetKey("sasl.mechanisms", config.Mechanism)
	}

	if config.SecurityProtocol != "" {
		kafkaConfig.SetKey("security.protocol", config.SecurityProtocol)
	}

	producer, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func newConsumer(config Config, group string, offsetOption int) (*kafka.Consumer, error) {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.Brokers,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	}

	if config.Username != "" && config.Password != "" {
		kafkaConfig.SetKey("sasl.username", config.Username)
		kafkaConfig.SetKey("sasl.password", config.Password)
		kafkaConfig.SetKey("sasl.mechanisms", config.Mechanism)
	}

	if config.SecurityProtocol != "" {
		kafkaConfig.SetKey("security.protocol", config.SecurityProtocol)
	}

	switch offsetOption {
	case OffsetFromLatest:
		kafkaConfig.SetKey("auto.offset.reset", "latest")
	case OffsetFromEarliest:
		kafkaConfig.SetKey("auto.offset.reset", "earliest")
	}

	consumer, err := kafka.NewConsumer(kafkaConfig)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func (r *client) IsConnected() bool {
	return r.producer != nil
}

func (r *client) Publish(topic string, message interface{}) error {
	bData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          bData,
	}

	return r.producer.Produce(msg, nil)
}

func (r *client) Subscribe(topics []string, offsetOption int, consumerGroup string) (<-chan *Message, error) {
	messages := make(chan *Message)
	consumer, err := newConsumer(Config{Brokers: r.brokers}, consumerGroup, offsetOption)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			ev := consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				messages <- &Message{
					Topic:     *e.TopicPartition.Topic,
					Partition: e.TopicPartition.Partition,
					Offset:    int64(e.TopicPartition.Offset),
					Key:       e.Key,
					Value:     e.Value,
					Timestamp: e.Timestamp,
				}
			case kafka.Error:
				log.Printf("Error: %v\n", e)
			}
		}
	}()

	r.consumer = consumer

	return messages, nil
}

func (r *client) Close() error {
	if r.producer != nil {
		r.producer.Close()
	}
	if r.consumer != nil {
		r.consumer.Close()
	}
	return nil
}
