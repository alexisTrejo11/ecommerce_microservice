package config

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

const (
	QueueUpdateCourseRating = "update_course_rating"
)

func ConnectRabbitMQ() (*amqp.Connection, error) {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		rabbitMQURL = "amqp://guest:guest@review-rabbitmq:5672/"
	}

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	return conn, nil
}

func DeclareQueues(channel *amqp.Channel) {
	queueName := os.Getenv("RABBITMQ_QUEUE_NAME")
	_, err := channel.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Auto-Delete
		false, // Exclusive
		false, // No-Wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue %s: %v", queueName, err)
	}

	_, err = channel.QueueDeclare(
		QueueUpdateCourseRating,
		true,  // Durable
		false, // Auto-Delete
		false, // Exclusive
		false, // No-Wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue %s: %v", QueueUpdateCourseRating, err)
	}
}
