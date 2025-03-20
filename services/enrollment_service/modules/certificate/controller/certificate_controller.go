package controller

import (
	"context"

	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
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
	logging.LogIncomingRequest(c, "get_my_certificates")

	userID, err := cc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_enrollment_id")
	}

	certificates, err := cc.service.GetCertificateByUserID(context.Background(), userID)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_my_certificates", userID.String())
	}

	logging.LogSuccess("get_my_certificates", "User Certificate Succesfully Retrieved", map[string]interface{}{
		"user_id": userID,
	})

	return response.OK(c, "User Certificate Succesfully Retrieved", certificates)
}

func (cc *CertifcateController) GetCertificateByEnrollment(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_certificate_by_enrollment")

	enrollmentID, err := response.GetUUIDParam(c, "enrollment_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_enrollment_id")
	}

	certificate, err := cc.service.GetCertificateByEnrollment(context.Background(), enrollmentID)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_certificate_by_enrollment", enrollmentID.String())
	}

	logging.LogSuccess("get_certificate_by_enrollment", "User Certificate Succesfully Retrieved", map[string]interface{}{
		"enrollment_id": enrollmentID,
	})

	return response.OK(c, "Certificate Succesfully Retrieved", certificate)
}
