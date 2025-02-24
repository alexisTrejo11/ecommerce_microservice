package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type EmailMessage struct {
	UserID            string `json:"user_id"`
	VerificationToken string `json:"verification_token"`
}

func ConnectRabbitMQ() (*amqp.Connection, error) {
	return amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
}

func PublishMessage(ctx context.Context, conn *amqp.Connection, queueName string, message EmailMessage) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Auto-Delete
		false, // Exclusive
		false, // No-Wait
		nil,
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	log.Println("Message send to queue:", queueName)
	return nil
}

type EmailConsumer struct {
	emailUseCase input.EmailUseCase
}

func NewEmailConsumer(emailUseCase input.EmailUseCase) *EmailConsumer {
	return &EmailConsumer{emailUseCase: emailUseCase}
}

func (uc *EmailConsumer) ConsumeEmail() {
	conn, err := ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email_queue",
		true,  // Durable
		false, // Auto-Delete
		false, // Exclusive
		false, // No-Wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,  // Auto-Acknowledge
		false, // Exclusive
		false, // No-Local
		false, // No-Wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	fmt.Println("Worker listening for email messages...")

	go func() {
		for msg := range msgs {
			var emailMessage EmailMessage
			err := json.Unmarshal(msg.Body, &emailMessage)
			if err != nil {
				log.Println("Error decoding message:", err)
				continue
			}

			fmt.Printf("Sending verification email to UserID: %s\n", emailMessage.UserID)

			userId, _ := uuid.Parse(emailMessage.UserID)
			err = uc.emailUseCase.SendVerificationEmail(context.Background(), userId, emailMessage.VerificationToken)
			if err != nil {
				log.Println("Error sending email:", err)
			} else {
				log.Printf("Verification email successfully sent to UserID: %s\n", emailMessage.UserID)
			}
		}
	}()

	select {}
}
