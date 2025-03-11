package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/input/v1/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func LessonRoutes(app *fiber.App, lessonHanlders handlers.LessonHandler) {
	path := app.Group("v1/api/lessons")
	path.Get("/:id", lessonHanlders.GetLessonById)
	path.Post("", lessonHanlders.CreateLesson)
	path.Put("/:id", lessonHanlders.UpdateLesson)
	path.Delete("/:id", lessonHanlders.DeleteLession)
}
