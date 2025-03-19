package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/controller"
	"github.com/gofiber/fiber/v2"
)

func EnrollmentsRoutes(app *fiber.App, commandController controller.EnrollmentComandController, queryController controller.EnrollmentQueryController) {
	path := app.Group("/v1/api/enrollments")

	path.Get("/:course_id", queryController.GetCourseEnrollments)
	path.Get("/:enrollent_id", queryController.GetEnrollmentByID)
	path.Get("/:user_id/:course_id", queryController.GetEnrollmentByUserAndCourse)
	path.Get("/my", queryController.GetUserEnrollments)

	path.Put("/:enrollent_id/complete", commandController.CompleteCourse)
	path.Put("/:enrollent_id/cancel", commandController.CancellEnrollment)
	path.Post("/:user_id/:course_id", commandController.EnrollUserInCourse)

}
