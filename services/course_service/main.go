package main

import (
	"log"
	"os"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.GORMConfig()
	config.InitRedis()

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Course Service")
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
