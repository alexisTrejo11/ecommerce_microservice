package routes

import (
	handler "github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/ports/input/handler/v1/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func NotificationRoutes(app *fiber.App, notificaionHandler *handler.NotificationHandler) {
	// Home
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"Home": "Welcome to notification service"})
	})

	path := app.Group("/v1/api/notifications")

	path.Get("user/:user_id", notificaionHandler.GetNotificationByUserId)
	path.Get("/:id", notificaionHandler.GetNotificationById)
	path.Delete("/:id", notificaionHandler.CancelNotification)
}
