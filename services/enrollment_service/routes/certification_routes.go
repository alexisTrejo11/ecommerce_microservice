package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/controller"
	"github.com/gofiber/fiber/v2"
)

func CerticationRoutes(app *fiber.App, controller controller.CertifcateController) {
	path := app.Group("/v1/api/certifications")
	path.Get("/:enrollment_id", controller.GetCertificateByEnrollment)
}

func UserCerticationRoutes(app *fiber.App, controller controller.CertifcateController) {
	path := app.Group("/v1/api/certifications")

	path.Get("users/my", controller.GetMyCertificates)
}
