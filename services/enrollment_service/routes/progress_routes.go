package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/controller"
	"github.com/gofiber/fiber/v2"
)

func ProgressRoutes(app *fiber.App, controller controller.ProgressController) {
	path := app.Group("/v1/api/enrollments/courses/progress")

	path.Get("/users/my", controller.GetMyCourseProgress)
	// Ad by Id
	path.Put("/:lesson_id/complete", controller.MarkLessonComplete)
	path.Put("/:lesson_id/uncomplete", controller.MarkLessonIncomplete)

}
