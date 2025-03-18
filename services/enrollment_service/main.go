package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Enrollment struct {
}

func main() {
	app := fiber.New()

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Welcome To Enrollment Service")
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
