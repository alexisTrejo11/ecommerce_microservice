package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/input/v1/http/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func CartRoutes(app *fiber.App, cartHandler handlers.CartHandler) {
	path := app.Group("/api/v1")

	path.Post("/carts/:userId", cartHandler.InitCart)
	path.Get("/carts/:userId", cartHandler.GetCartByUserId)
	path.Get("/carts/id/:id", cartHandler.GetCartById)
	path.Get("/carts/id/:id", cartHandler.DeleteCart)
}

func UserCartRoutes(app *fiber.App, userCartHandler handlers.UserCartHandler) {
	path := app.Group("/api/v1/user")

	path.Get("/carts", userCartHandler.GetMyCart)
	path.Post("/carts/items", userCartHandler.AddItems)
	path.Delete("/carts/item", userCartHandler.RemoveItems)
	path.Post("/cart/buy", userCartHandler.RemoveItems)
	path.Post("/cart/buy-product", userCartHandler.RemoveItems)
}
