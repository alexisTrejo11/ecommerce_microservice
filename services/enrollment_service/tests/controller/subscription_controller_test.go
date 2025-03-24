package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/controller"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/middleware"
	mocks "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/tests/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteSubscription_Success(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	subscriptionID := uuid.New()

	mockService.On("DeleteSubscription", mock.Anything, subscriptionID).Return(nil)

	app.Delete("/v1/api/subscriptions/:subscription_id", func(c *fiber.Ctx) error {

		c.Request().Header.Set("Authorization", "Bearer test-token")
		return subscriptionController.DeleteSubscription(c)
	})

	req := httptest.NewRequest("DELETE", "/v1/api/subscriptions/"+subscriptionID.String(), nil)
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestDeleteSubscription_InvalidSubscriptionID(t *testing.T) {

	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	app.Delete("/v1/api/subscriptions/:subscription_id", func(c *fiber.Ctx) error {

		c.Request().Header.Set("Authorization", "Bearer test-token")
		return subscriptionController.DeleteSubscription(c)
	})

	req := httptest.NewRequest("DELETE", "/v1/api/subscriptions/invalid-uuid", nil)
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, 400, resp.StatusCode)
}

func TestDeleteSubscription_ServiceError(t *testing.T) {

	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &jwt.JWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	subscriptionID := uuid.New()

	mockService.On("DeleteSubscription", mock.Anything, subscriptionID).Return(errors.New("service error"))

	app.Delete("/v1/api/subscriptions/:subscription_id", func(c *fiber.Ctx) error {

		c.Request().Header.Set("Authorization", "Bearer test-token")
		return subscriptionController.DeleteSubscription(c)
	})

	req := httptest.NewRequest("DELETE", "/v1/api/subscriptions/"+subscriptionID.String(), nil)
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, 500, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestCreateSubscription_Success(t *testing.T) {
	// Arrange
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	requestBody := dtos.SubscriptionInsertDTO{
		UserID:    uuid.MustParse("99eab094-d405-4a43-8894-7f1c83f162db"),
		PlanName:  "premuim subscription",
		Status:    "ACTIVE",
		Type:      "MONTHLY",
		PaymentID: uuid.MustParse("99eab094-d405-4a43-8894-7f1c83f162db"),
	}

	mockResponse := &dtos.SubscriptionDTO{
		ID:        uuid.New(),
		UserID:    requestBody.UserID,
		PlanName:  requestBody.PlanName,
		Status:    requestBody.Status,
		Type:      requestBody.Type,
		PaymentID: requestBody.PaymentID,
	}

	mockService.On("CreateSubscription", mock.Anything, mock.MatchedBy(func(dto dtos.SubscriptionInsertDTO) bool {
		return dto.UserID == requestBody.UserID &&
			dto.PlanName == requestBody.PlanName &&
			dto.Status == requestBody.Status &&
			dto.Type == requestBody.Type &&
			dto.PaymentID == requestBody.PaymentID
	})).Return(mockResponse, nil)

	app.Post("/v1/api/subscriptions", func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer test-token")
		return subscriptionController.CreateSubscription(c)
	})

	reqBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/v1/api/subscriptions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	mockService.AssertExpectations(t)
	mockJWTManager.AssertExpectations(t)
}

func TestCreateSubscription_InvalidBody(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	app.Post("/v1/api/subscriptions", func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer test-token")
		return subscriptionController.CreateSubscription(c)
	})

	req := httptest.NewRequest("POST", "/v1/api/subscriptions", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestCreateSubscription_ValidationError(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	requestBody := dtos.SubscriptionInsertDTO{
		UserID:    uuid.Nil, // Invalid UUID
		PlanName:  "",
		Status:    "INVALID_STATUS",
		Type:      "INVALID_TYPE",
		PaymentID: uuid.Nil,
	}

	mockJWTManager.On("ExtractUserID", "test-token").Return(uuid.New(), nil)

	app.Post("/v1/api/subscriptions", func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer test-token")
		return subscriptionController.CreateSubscription(c)
	})

	reqBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/v1/api/subscriptions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestCreateSubscription_ServiceError(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	requestBody := dtos.SubscriptionInsertDTO{
		UserID:    uuid.MustParse("99eab094-d405-4a43-8894-7f1c83f162db"),
		PlanName:  "premium",
		Status:    "ACTIVE",
		Type:      "MONTHLY",
		PaymentID: uuid.New(),
	}

	mockService.On("CreateSubscription", mock.Anything, mock.Anything).Return(nil, errors.New("service error"))
	mockJWTManager.On("ExtractUserID", "test-token").Return(requestBody.UserID, nil)

	app.Post("/v1/api/subscriptions", func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer test-token")
		return subscriptionController.CreateSubscription(c)
	})

	reqBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/v1/api/subscriptions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestGetMySubscription_Success(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	userID := uuid.MustParse("ff8c6bc9-2a2c-4ab9-b65f-88deb761b5de")
	mockSubscription := &dtos.SubscriptionDTO{
		ID:        uuid.New(),
		UserID:    userID,
		PlanName:  "premium",
		Status:    "ACTIVE",
		Type:      "MONTHLY",
		PaymentID: uuid.New(),
	}

	// Configurar mocks
	mockService.On("GetSubscriptionByUser", mock.Anything, userID).Return(mockSubscription, nil)
	mockJWTManager.On("GetUserIDFromToken", mock.Anything).Return(userID, nil)

	// Aplicar el middleware de autenticación
	app.Use(middleware.JWTAuthMiddleware(mockJWTManager))

	// Definir la ruta
	app.Get("/v1/api/subscriptions/me", func(c *fiber.Ctx) error {
		return subscriptionController.GetMySubscription(c)
	})

	// Crear la solicitud
	req := httptest.NewRequest("GET", "/v1/api/subscriptions/me", nil)
	req.Header.Set("Authorization", "Bearer test-token")

	// Ejecutar la solicitud
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetMySubscription_Unauthorized(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	mockJWTManager.On("ExtractUserID", "").Return(uuid.Nil, errors.New("invalid token"))

	app.Get("/v1/api/subscriptions/me", func(c *fiber.Ctx) error {
		return subscriptionController.GetMySubscription(c)
	})

	app.Use(middleware.JWTAuthMiddleware(mockJWTManager))
	req := httptest.NewRequest("GET", "/v1/api/subscriptions/me", nil)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestGetMySubscription_ServiceError(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	userID := uuid.MustParse("ff8c6bc9-2a2c-4ab9-b65f-88deb761b5de")

	mockService.On("GetSubscriptionByUser", mock.Anything, userID).Return(nil, errors.New("DATABASE_ERROR"))
	mockJWTManager.On("ExtractUserID", "test-token").Return(userID, nil)

	app.Use(middleware.JWTAuthMiddleware(mockJWTManager))
	app.Get("/v1/api/subscriptions/me", func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer test-token")
		return subscriptionController.GetMySubscription(c)
	})

	req := httptest.NewRequest("GET", "/v1/api/subscriptions/me", nil)
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestCancelMySubscription_Success(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	userID := uuid.MustParse("ff8c6bc9-2a2c-4ab9-b65f-88deb761b5de")
	subscriptionID := uuid.New()

	mockService.On("CancelSubscription", mock.Anything, userID, subscriptionID).Return(nil)
	mockJWTManager.On("ExtractUserID", "test-token").Return(userID, nil)

	app.Use(middleware.JWTAuthMiddleware(mockJWTManager))
	app.Post("/v1/api/subscriptions/cancel/:lesson_id", func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer test-token")
		c.Params("lesson_id", subscriptionID.String())
		return subscriptionController.CancelMySubscription(c)
	})

	req := httptest.NewRequest("POST", "/v1/api/subscriptions/cancel/"+subscriptionID.String(), nil)
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestCancelMySubscription_InvalidID(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	mockJWTManager.On("ExtractUserID", "test-token").Return(uuid.New(), nil)

	app.Use(middleware.JWTAuthMiddleware(mockJWTManager))
	app.Post("/v1/api/subscriptions/cancel/:lesson_id", func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer test-token")
		c.Params("lesson_id", "invalid-uuid")
		return subscriptionController.CancelMySubscription(c)
	})

	req := httptest.NewRequest("POST", "/v1/api/subscriptions/cancel/invalid-uuid", nil)
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestCancelMySubscription_Unauthorized(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	mockJWTManager.On("ExtractUserID", "").Return(uuid.Nil, errors.New("invalid token"))

	app.Post("/v1/api/subscriptions/cancel/:lesson_id", func(c *fiber.Ctx) error {
		return subscriptionController.CancelMySubscription(c)
	})

	req := httptest.NewRequest("POST", "/v1/api/subscriptions/cancel/some-id", nil)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestCancelMySubscription_ServiceError(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.MockSubscriptionService)
	mockJWTManager := &mocks.MockJWTManager{}

	subscriptionController := controller.NewSubscriptionController(mockService, mockJWTManager)

	userID := uuid.MustParse("ff8c6bc9-2a2c-4ab9-b65f-88deb761b5de")
	subscriptionID := uuid.New()

	mockService.On("CancelSubscription", mock.Anything, userID, subscriptionID).Return(errors.New("service error"))
	mockJWTManager.On("ExtractUserID", "test-token").Return(userID, nil)

	app.Use(middleware.JWTAuthMiddleware(mockJWTManager))
	app.Post("/v1/api/subscriptions/cancel/:lesson_id", func(c *fiber.Ctx) error {
		c.Request().Header.Set("Authorization", "Bearer test-token")
		c.Params("lesson_id", subscriptionID.String())
		return subscriptionController.CancelMySubscription(c)
	})

	req := httptest.NewRequest("POST", "/v1/api/subscriptions/cancel/"+subscriptionID.String(), nil)
	req.Header.Set("Authorization", "Bearer test-token")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
}
