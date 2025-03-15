package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	rabbitmq "github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/messaging"
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func ConnectRabbitMQ() (*amqp.Connection, error) {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		rabbitMQURL = "amqp://guest:guest@rabbitmq:5672/"
	}

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	return conn, nil
}

func NewRabbitMQClient(conn *amqp.Connection) (*RabbitMQClient, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &RabbitMQClient{conn: conn, channel: channel}, nil
}

func (r *RabbitMQClient) ReceiveMessage(queueName string, timeout time.Duration) ([]byte, string, error) {
	msgs, err := r.channel.Consume(
		queueName,
		"",    // Consumer (vacío para que RabbitMQ asigne uno)
		false, // AutoAck (false para manejar confirmación manual)
		false, // Exclusive
		false, // NoLocal
		false, // NoWait
		nil,   // Args
	)
	if err != nil {
		return nil, "", fmt.Errorf("failed to consume messages: %w", err)
	}

	messageCh := make(chan amqp.Delivery)
	go func() {
		for msg := range msgs {
			messageCh <- msg
			return
		}
	}()

	select {
	case msg := <-messageCh:
		return msg.Body, fmt.Sprintf("%d", msg.DeliveryTag), nil
	case <-time.After(timeout):
		return nil, "", errors.New("timeout waiting for message")
	}
}

func (r *RabbitMQClient) DeleteMessage(queueName string, receiptHandle string) error {
	deliveryTag, err := parseDeliveryTag(receiptHandle)
	if err != nil {
		return fmt.Errorf("invalid receipt handle: %w", err)
	}
	return r.channel.Ack(deliveryTag, false)
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

func NewReceiverNotificationQueue(notificationUseCase input.NotificationUseCase, queueClient *RabbitMQClient) *rabbitmq.QueueReceiver {
	queueName := os.Getenv("RABBITMQ_QUEUE_NAME")
	timeoutStr := os.Getenv("QUEUE_TIMEOUT_SECONDS")

	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Fatalf("Invalid timeout value: %v", err)
	}

	queueTimeout := time.Duration(timeout) * time.Second

	return rabbitmq.NewQueueReceiver(queueClient, queueName, queueTimeout, notificationUseCase)
}
