package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/config"
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQClient(conn *amqp.Connection) (*RabbitMQClient, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	config.DeclareQueues(channel)

	return &RabbitMQClient{conn: conn, channel: channel}, nil
}

func (r *RabbitMQClient) PublishMessage(queueName string, payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling message: %w", err)
	}

	err = r.channel.Publish(
		"",        // Exchange (vacío para usar default)
		queueName, // Routing key (nombre de la cola)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent, // Mensaje persistente
			Timestamp:    time.Now(),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message to queue %s: %w", queueName, err)
	}

	return nil
}

func (r *RabbitMQClient) Close() {
	if err := r.channel.Close(); err != nil {
		log.Println("Error closing channel:", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Println("Error closing connection:", err)
	}
}

func parseDeliveryTag(receiptHandle string) (uint64, error) {
	var tag uint64
	_, err := fmt.Sscanf(receiptHandle, "%d", &tag)
	return tag, err
}
