package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router, authHandler *handlers.AuthHandler) {
	r.Get("/home", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Welcome": "Welcome to User Service",
		})
	})

	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)

}
