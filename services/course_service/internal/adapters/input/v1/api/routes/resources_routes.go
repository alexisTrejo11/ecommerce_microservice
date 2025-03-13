package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/input/v1/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func ResourceRoutes(app *fiber.App, resourceHanlders handlers.ResourceHandler) {
	path := app.Group("v1/api/resources")
	path.Get("/:id", resourceHanlders.GetResourceById)
	path.Get("/lesson/:lesson_id", resourceHanlders.GetResourcesByLessonId)
	path.Post("", resourceHanlders.CreateResource)
	path.Put("/:id", resourceHanlders.UpdateResource)
	path.Delete("/:id", resourceHanlders.DeleteResource)
}
