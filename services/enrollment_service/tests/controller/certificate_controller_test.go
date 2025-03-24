package controller_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/controller"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/middleware"
	mocks "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/tests/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMyCertificates_Success(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockCertificateService)
	mockJWTManager := &mocks.MockJWTManager{}

	certificateController := controller.NewCertificateController(mockService, mockJWTManager)

	userID := uuid.MustParse("ff8c6bc9-2a2c-4ab9-b65f-88deb761b5de")
	mockCertificates := createMockCertificates()

	mockService.On("GetCertificateByUserID", mock.Anything, userID).Return(mockCertificates, nil)

	app.Use(middleware.JWTAuthMiddleware(mockJWTManager))
	app.Use(func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer test-token")
		return c.Next()
	})

	app.Get("/v1/api/certificates/my", func(c *fiber.Ctx) error {
		return certificateController.GetMyCertificates(c)
	})

	req := httptest.NewRequest("GET", "/v1/api/certificates/my", nil)
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	mockJWTManager.AssertExpectations(t)
	mockService.AssertExpectations(t)
}

func TestGetMyCertificates_Unauthorized(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockCertificateService)
	mockJWTManager := &mocks.MockJWTManager{}

	certificateController := controller.NewCertificateController(mockService, mockJWTManager)

	app.Use(func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer invalid-token")
		return c.Next()
	})

	app.Get("/v1/api/certificates/my", func(c *fiber.Ctx) error {
		return certificateController.GetMyCertificates(c)
	})

	req := httptest.NewRequest("GET", "/v1/api/certificates/my", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	mockJWTManager.AssertExpectations(t)
	mockService.AssertNotCalled(t, "GetCertificateByUserID", mock.Anything, mock.Anything)
}

func createMockCertificates() *[]dtos.CertificateDTO {
	now := time.Now()
	expiresAt := now.AddDate(1, 0, 0)

	return &[]dtos.CertificateDTO{
		{
			ID:             uuid.New(),
			EnrollmentID:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			IssuedAt:       now,
			CertificateURL: "https://example.com/certificates/12345",
			ExpiresAt:      &expiresAt,
		},
		{
			ID:             uuid.New(),
			EnrollmentID:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174001"),
			IssuedAt:       now.AddDate(0, 6, 0), // 6 meses después
			CertificateURL: "https://example.com/certificates/67890",
			ExpiresAt:      nil, // Certificado sin fecha de expiración
		},
	}
}
