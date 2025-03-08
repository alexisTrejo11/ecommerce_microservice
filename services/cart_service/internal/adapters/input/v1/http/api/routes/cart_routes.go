package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/input/v1/http/api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CartRoutes(app *fiber.App, cartHandler handlers.CartHandler) {
	path := app.Group("/v1/api")

	path.Post("/carts/user/:userId", cartHandler.InitCart)
	path.Get("/carts/user/:userId", cartHandler.GetCartByUserId)
	path.Get("/carts/:id", cartHandler.GetCartById)
	path.Delete("/carts/:id", cartHandler.DeleteCart)

	path.Get("/uuid", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(uuid.New())
	})
}

func UserCartRoutes(app *fiber.App, userCartHandler handlers.UserCartHandler) {
	path := app.Group("/v1/api/user")

	path.Get("/carts/:userId", userCartHandler.GetMyCart)
	path.Post("/carts/items/:id", userCartHandler.AddItems)
	path.Delete("/carts/items/:id", userCartHandler.RemoveItems)
	path.Post("/carts/buy", userCartHandler.RemoveItems)
	path.Post("/carts/buy-product", userCartHandler.RemoveItems)
}
