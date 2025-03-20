package controller

import (
	"context"

	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	su_service "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
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
	logging.LogIncomingRequest(c, "delete_subscription")

	subscriptionID, err := response.GetUUIDParam(c, "subscription_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_subscription_id")
	}

	err = sc.service.DeleteSubscription(context.Background(), subscriptionID)
	if err != nil {
		return response.HandleApplicationError(c, err, "delete_subscription", subscriptionID.String())
	}

	logging.LogSuccess("delete_subscription", "Subscription Successfully Deleted.", map[string]interface{}{
		"subscription_id": subscriptionID,
	})

	return response.OK(c, "Subscription Successfully Deleted.", nil)
}

func (sc *SubscriptionController) ChangeMySubscriptionType(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "change_my_subscription_type")

	subscriptionID, err := response.GetUUIDParam(c, "user_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_lesson_id")
	}

	typeSTR := c.Params("sub_type")
	if typeSTR == "" {
		return response.BadRequest(c, "Sub Type Can't be Empty", "invalid_param")
	}

	subType := subscription.SubscriptionType(typeSTR)
	if !subType.IsValid() {
		return response.BadRequest(c, "Invalid Subscription Type", "invalid_param")
	}

	err = sc.service.UpdateSubscriptionType(context.Background(), subscriptionID, subType)
	if err != nil {
		return response.HandleApplicationError(c, err, "change_my_subscription_type", subscriptionID.String())
	}

	logging.LogSuccess("change_my_subscription_type", "Subscription Successfully Upated.", map[string]interface{}{
		"subscription_id": subscriptionID,
	})

	return response.OK(c, "Subscription Successfully Upated.", nil)
}

func (sc *SubscriptionController) CreateSubscription(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "create_subscription")

	var insertDTO dtos.SubscriptionInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, err.Error(), "invalid_request_body")
	}

	if err := sc.validator.Struct(&insertDTO); err != nil {
		return response.BadRequest(c, err.Error(), "invalid_request_data")
	}

	subscription, err := sc.service.CreateSubscription(context.TODO(), insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, "create_subscription", insertDTO.UserID.String())
	}

	logging.LogSuccess("create_subscription", "Subscription Successfully Created", map[string]interface{}{
		"subscription_id": subscription.ID,
	})

	return response.Created(c, "Subscription Successfully Created", subscription)
}

// User
func (sc *SubscriptionController) GetMySubscription(c *fiber.Ctx) error {
	userID, err := sc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	logging.LogIncomingRequest(c, "get_my_subscription")

	subscription, err := sc.service.GetSubscriptionByUser(context.Background(), userID)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_my_subscription", userID.String())
	}

	logging.LogSuccess("get_my_subscription", "Subscription Successfully Retrieved.", map[string]interface{}{
		"subscription_id": subscription.ID,
	})

	return response.Created(c, "Subscription Successfully Retrieved", subscription)
}

func (sc *SubscriptionController) CancelMySubscription(c *fiber.Ctx) error {
	userID, err := sc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	logging.LogIncomingRequest(c, "cancel_my_subscription")

	subscriptionID, err := response.GetUUIDParam(c, "lesson_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_lesson_id")
	}

	err = sc.service.CancelSubscription(context.Background(), userID, subscriptionID)
	if err != nil {
		response.BadRequest(c, err.Error(), "invalid_input")
	}

	logging.LogSuccess("cancel_my_subscription", "Subscription Successfully Cancelled.", map[string]interface{}{
		"subscription_id": subscriptionID,
	})

	return response.OK(c, "Subscription Successfully Cancelled.", nil)
}
