package controller

import (
	"context"

	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/response"
	"github.com/gofiber/fiber/v2"
)

type CertifcateController struct {
	service    services.CertificateService
	jwtManager jwt.JWTManager
}

func NewCertificateController(service services.CertificateService, jwtManager jwt.JWTManager) *CertifcateController {
	return &CertifcateController{
		service:    service,
		jwtManager: jwtManager,
	}
}

func (cc *CertifcateController) GetMyCertificates(c *fiber.Ctx) error {
	userID, err := cc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_enrollment_id")
	}

	certificates, err := cc.service.GetCertificateByUserID(context.Background(), userID)
	if err != nil {
		return response.BadRequest(c, err.Error(), "error")
	}

	return response.OK(c, "User Certificate Succesfully Retrieved", certificates)
}

func (cc *CertifcateController) GetCertificateByEnrollment(c *fiber.Ctx) error {
	enrollmentID, err := response.GetUUIDParam(c, "enrollment_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_enrollment_id")
	}

	certificate, err := cc.service.GetCertificateByEnrollment(context.Background(), enrollmentID)
	if err != nil {
		return response.BadRequest(c, err.Error(), "error")
	}

	return response.OK(c, "Certificate Succesfully Retrieved", certificate)
}
