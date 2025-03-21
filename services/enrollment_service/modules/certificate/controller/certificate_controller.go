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
	logging.LogIncomingRequest(c, KeyGetMyCertificates)

	userID, err := cc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidEnrollmentID)
	}

	certificates, err := cc.service.GetCertificateByUserID(context.Background(), userID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyGetMyCertificates, userID.String())
	}

	logging.LogSuccess(KeyGetMyCertificates, MsgUserCertificateRetrieved, map[string]interface{}{
		"user_id": userID,
	})

	return response.OK(c, MsgUserCertificateRetrieved, certificates)
}

func (cc *CertifcateController) GetCertificateByEnrollment(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, KeyGetCertificateByEnrollment)

	enrollmentID, err := response.GetUUIDParam(c, "enrollment_id", KeyGetCertificateByEnrollment)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidEnrollmentID)
	}

	certificate, err := cc.service.GetCertificateByEnrollment(context.Background(), enrollmentID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyGetCertificateByEnrollment, enrollmentID.String())
	}

	logging.LogSuccess(KeyGetCertificateByEnrollment, MsgCertificateRetrieved, map[string]interface{}{
		"enrollment_id": enrollmentID,
	})

	return response.OK(c, MsgCertificateRetrieved, certificate)
}
