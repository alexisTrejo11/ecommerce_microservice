package controller

import (
	"context"

	suscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/model"
	su_service "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SubscriptionController struct {
	service    su_service.SubscriptionService
	jwtManager jwt.JWTManager
	validator  *validator.Validate
}

func NewSubscriptionController(service su_service.SubscriptionService, jwtManager jwt.JWTManager) *SubscriptionController {
	return &SubscriptionController{
		service:    service,
		jwtManager: jwtManager,
		validator:  validator.New(),
	}
}

func (sc *SubscriptionController) DeleteSubscription(c *fiber.Ctx) error {
	subscriptionID, err := response.GetUUIDParam(c, "subscription_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_subscription_id")
	}

	err = sc.service.DeleteSubscription(context.Background(), subscriptionID)
	if err != nil {
		response.BadRequest(c, err.Error(), "invalid_input")
	}

	return response.OK(c, "Subscription Successfully Deleted.", nil)
}

func (sc *SubscriptionController) UpdateSubscriptionType(c *fiber.Ctx) error {
	subscriptionID, err := response.GetUUIDParam(c, "lesson_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_lesson_id")
	}

	typeSTR := c.Params("sub_type")
	if typeSTR == "" {
		return response.BadRequest(c, "Sub Type Can't be Empty", "invalid_param")
	}

	subType := suscription.SubscriptionType(typeSTR)
	if !subType.IsValid() {
		return response.BadRequest(c, "Invalid Subscription Type", "invalid_param")
	}

	err = sc.service.UpdateSubscriptionType(context.Background(), subscriptionID, subType)
	if err != nil {
		response.BadRequest(c, err.Error(), "invalid_input")
	}

	return response.OK(c, "Subscription Successfully Deleted.", nil)
}

func (sc *SubscriptionController) CreateSubscription(c *fiber.Ctx) error {
	var insertDTO dtos.SubscriptionInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		response.BadRequest(c, err.Error(), "invalid_request_body")
	}

	if err := sc.validator.Struct(&insertDTO); err != nil {
		response.BadRequest(c, err.Error(), "invalid_request_data")
	}

	subscription, err := sc.service.CreateSubscription(context.Background(), insertDTO)
	if err != nil {
		response.BadRequest(c, err.Error(), "invalid_input")
	}

	return response.Created(c, "Subscription Successfully Created", subscription)
}

// User
func (sc *SubscriptionController) GetMySubscription(c *fiber.Ctx) error {
	userID, err := sc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	subscription, err := sc.service.GetSubscriptionByUser(context.Background(), userID)
	if err != nil {
		response.BadRequest(c, err.Error(), "invalid_input")
	}

	return response.Created(c, "Subscription Successfully Created", subscription)
}

func (sc *SubscriptionController) CancelMySubscription(c *fiber.Ctx) error {
	userID, err := sc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	subscriptionID, err := response.GetUUIDParam(c, "lesson_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_lesson_id")
	}

	err = sc.service.CancelSubscription(context.Background(), userID, subscriptionID)
	if err != nil {
		response.BadRequest(c, err.Error(), "invalid_input")
	}

	return response.OK(c, "Subscription Successfully Cancelled.", nil)
}
