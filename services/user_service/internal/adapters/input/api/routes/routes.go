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

	path := r.Group("v1/api")
	path.Post("/register", authHandler.Register)
	path.Post("/login", authHandler.Login)
}

func UserRoutes(r fiber.Router, addresHandler *handlers.UserAddressHandler) {
	path := r.Group("v1/api/users/address")
	path.Get("", addresHandler.MyAddresses)
	path.Post("", addresHandler.AddAddress)
	path.Put("/:id", addresHandler.UpdateMyAddress)
	path.Delete("/:id", addresHandler.DeleteAddress)
}
