package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/input/v1/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func CourseRoutes(app *fiber.App, courseHanlders handlers.CourseHandler) {
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Course Service")
	})

	path := app.Group("v1/api/courses")
	path.Get("/:id", courseHanlders.GetCourseById)
	path.Post("", courseHanlders.CreateHandler)
	path.Put("/:id", courseHanlders.UpdateHandler)
	path.Delete("/:id", courseHanlders.DeleteLession)
}
