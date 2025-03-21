package controller

import (
	"context"

	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/response"
	"github.com/gofiber/fiber/v2"
)

// CertifcateController handles certificate-related HTTP requests
type CertifcateController struct {
	service    services.CertificateService
	jwtManager jwt.JWTManager
}

// NewCertificateController creates a new CertifcateController instance
func NewCertificateController(service services.CertificateService, jwtManager jwt.JWTManager) *CertifcateController {
	return &CertifcateController{
		service:    service,
		jwtManager: jwtManager,
	}
}

// GetMyCertificates godoc
// @Summary      Get user's certificates
// @Description  Retrieve all certificates associated with the authenticated user
// @Tags         Certificates
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.ApiResponse{data=[]dtos.CertificateDTO} "Certificates successfully retrieved"
// @Failure      401  {object}  response.ApiResponse "Unauthorized"
// @Failure      404  {object}  response.ApiResponse "No certificates found"
// @Router       /v1/api/certificates/my [get]
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

// GetCertificateByEnrollment godoc
// @Summary      Get certificate by enrollment
// @Description  Retrieve a certificate associated with a specific enrollment ID
// @Tags         Certificates
// @Accept       json
// @Produce      json
// @Param        enrollment_id   path      string  true  "Enrollment ID"
// @Success      200  {object}  response.ApiResponse{data=dtos.CertificateDTO} "Certificate successfully retrieved"
// @Failure      400  {object}  response.ApiResponse "Invalid enrollment ID"
// @Failure      404  {object}  response.ApiResponse "Certificate not found"
// @Router       /v1/api/certificates/{enrollment_id} [get]
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
