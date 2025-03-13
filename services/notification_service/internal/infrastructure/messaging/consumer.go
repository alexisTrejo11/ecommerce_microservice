package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/config"
)

type EmailMessage struct {
	UserID            string `json:"user_id"`
	VerificationToken string `json:"verification_token"`
}

type NotificationConsumer struct {
	notificationUseCase input.NotificationUseCase
}

func NewNotificationConsumer(notificationUserCase input.NotificationUseCase) *NotificationConsumer {
	return &NotificationConsumer{notificationUseCase: notificationUserCase}
}

func (uc *NotificationConsumer) ConsumeEmail() {
	conn, err := config.ConnectRabbitMQ()
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

			/*

				userId, _ := uuid.Parse(emailMessage.UserID)
				err = uc.emailUseCase.SendVerificationEmail(context.Background(), userId, emailMessage.VerificationToken)
				if err != nil {
					log.Println("Error sending email:", err)
				} else {
					log.Printf("Verification email successfully sent to UserID: %s\n", emailMessage.UserID)
				}
			*/
		}
	}()

	select {}
}

func (uc *NotificationConsumer) ConsumeSMS() {

}
