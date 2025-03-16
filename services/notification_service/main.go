package main

import (
	"context"
	"log"
	"os"

	usecase "github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/use_case"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/config"
	handler "github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/ports/input/handler/v1/api/handlers"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/ports/input/handler/v1/api/routes"
	repository "github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/email"
	logging "github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/log"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/sms"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Router
	app := fiber.New()

	// Email Config
	emailConfig := config.NewEmailConfig()
	mailClient := email.NewMailClient(emailConfig)

	// RabbitMQ
	conn, err := config.ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	queueClient, err := config.NewRabbitMQClient(conn)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}

	// SMS Config
	smsConfig := config.NewSMSConfig()
	smsService := sms.NewSMSService(smsConfig)

	// Logger
	logging.InitLogger()

	// DB
	mongoClient := config.InitMongoClient()

	// Repository
	notificationRepository := repository.NewNotificationRepository(mongoClient)

	// Use Case
	emailClient := usecase.NewEmailUseCase(mailClient)
	notficationUseCase := usecase.NewNotificationUseCase(notificationRepository, emailClient, *smsService)

	// Notifaction Reciever Queue
	queueReceiver := config.NewReceiverNotificationQueue(notficationUseCase, queueClient)
	go queueReceiver.ReceiveNotification(context.Background())

	// Handler
	notificationHandler := handler.NewNotificationHandler(notficationUseCase)

	routes.NotificationRoutes(app, notificationHandler)

	// Run Server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	logging.Logger.Info().Msgf("Server running on port %s", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
