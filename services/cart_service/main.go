package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Cart Service")
	})

}
