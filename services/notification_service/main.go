package main

import (
	"log"
	"os"

	usecase "github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/use_case"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/config"
	repository "github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/ports/output"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Router
	app := fiber.New()

	// DB
	mongoClient := config.InitMongoClient()

	// Repository
	notificationRepository := repository.NewNotificationRepository(mongoClient)
	usecase.NewNotificationUseCase(notificationRepository)

	// Home
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"Home": "Welcome to notification service"})
	})

	// Run Server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
