package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/controller"
	"github.com/gofiber/fiber/v2"
)

func ProgressRoutes(app *fiber.App, controller controller.ProgressController) {
	app.Group("/v1/api/enrollments/progress")

	app.Get("/my", controller.GetMyCourseProgress)
	// Ad by Id
	app.Put("/:lesson_id/complete", controller.MarkLessonComplete)
	app.Put("/:lesson_id/uncomplete", controller.MarkLessonIncomplete)

}
