package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/controller"
	"github.com/gofiber/fiber/v2"
)

func EnrollmentsRoutes(app *fiber.App, commandController controller.EnrollmentComandController, queryController controller.EnrollmentQueryController) {
	path := app.Group("/v1/api/enrollments")

	path.Get("course/:course_id", queryController.GetCourseEnrollments)
	path.Get("/:enrollment_id", queryController.GetEnrollmentByID)
	path.Get("/user/:user_id/course/:course_id", queryController.GetEnrollmentByUserAndCourse)

	path.Put("/:enrollment_id/complete", commandController.CompleteCourse)
	path.Put("/:enrollment_id/cancel", commandController.CancellEnrollment)
	path.Post("/:user_id/:course_id", commandController.EnrollUserInCourse)

}

func UserEnrollmentsRoutes(app *fiber.App, commandController controller.EnrollmentComandController, queryController controller.EnrollmentQueryController) {
	path := app.Group("/v1/api/enrollments")
	path.Get("/user/my", queryController.GetMyEnrollments)
}
