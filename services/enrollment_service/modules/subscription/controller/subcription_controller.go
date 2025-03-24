package controller

import (
	"context"

	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	su_service "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/middleware"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// SubscriptionController handles operations related to subscriptions.
type SubscriptionController struct {
	service    su_service.SubscriptionService
	jwtManager jwt.TokenManager
	validator  *validator.Validate
}

// NewSubscriptionController creates a new instance of the SubscriptionController.
func NewSubscriptionController(service su_service.SubscriptionService, jwtManager jwt.TokenManager) *SubscriptionController {
	return &SubscriptionController{
		service:    service,
		jwtManager: jwtManager,
		validator:  validator.New(),
	}
}

// @Summary Delete Subscription
// @Description Deletes a subscription by its ID.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param subscription_id path string true "Subscription ID"
// @Success 200 {object} response.ApiResponse "Subscription deleted successfully"
// @Failure 400 {object} response.ApiResponse "Invalid subscription ID"
// @Failure 500 {object} response.ApiResponse "Internal server error"
// @Router /v1/api/subscriptions/{subscription_id} [delete]
func (sc *SubscriptionController) DeleteSubscription(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, KeyDeleteSubscription)

	subscriptionID, err := response.GetUUIDParam(c, "subscription_id", KeyDeleteSubscription)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidSubscriptionID)
	}

	err = sc.service.DeleteSubscription(context.Background(), subscriptionID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyDeleteSubscription, subscriptionID.String())
	}

	logging.LogSuccess(KeyDeleteSubscription, MsgSubscriptionDeleted, map[string]interface{}{
		"subscription_id": subscriptionID,
	})

	return response.OK(c, MsgSubscriptionDeleted, nil)
}

// @Summary Change My Subscription Type
// @Description Changes the subscription type for the authenticated user.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Param sub_type path string true "Subscription Type"
// @Success 200 {object} response.ApiResponse "Subscription type updated successfully"
// @Failure 400 {object} response.ApiResponse "Invalid subscription type or user ID"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Failure 500 {object} response.ApiResponse "Internal server error"
// @Router /v1/api/subscriptions/{user_id}/change/{sub_type} [post]
func (sc *SubscriptionController) ChangeMySubscriptionType(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, KeyChangeMySubscriptionType)

	subscriptionID, err := response.GetUUIDParam(c, "user_id", KeyChangeMySubscriptionType)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidSubscriptionID)
	}

	typeSTR := c.Params("sub_type")
	if typeSTR == "" {
		return response.BadRequest(c, "Sub Type Can't be Empty", MsgInvalidParam)
	}

	subType := subscription.SubscriptionType(typeSTR)
	if !subType.IsValid() {
		return response.BadRequest(c, "Invalid Subscription Type", MsgInvalidParam)
	}

	err = sc.service.UpdateSubscriptionType(context.Background(), subscriptionID, subType)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyChangeMySubscriptionType, subscriptionID.String())
	}

	logging.LogSuccess(KeyChangeMySubscriptionType, MsgSubscriptionUpdated, map[string]interface{}{
		"subscription_id": subscriptionID,
	})

	return response.OK(c, MsgSubscriptionUpdated, nil)
}

// @Summary Create Subscription
// @Description Creates a new subscription for a user.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body dtos.SubscriptionInsertDTO true "Subscription details"
// @Success 201 {object} response.ApiResponse "Subscription created successfully"
// @Failure 400 {object} response.ApiResponse "Invalid request data"
// @Failure 500 {object} response.ApiResponse "Internal server error"
// @Router /v1/api/subscriptions [post]
func (sc *SubscriptionController) CreateSubscription(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, KeyCreateSubscription)

	var insertDTO dtos.SubscriptionInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidRequestBody)
	}

	if err := sc.validator.Struct(&insertDTO); err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidRequestData)
	}

	subscription, err := sc.service.CreateSubscription(context.TODO(), insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyCreateSubscription, insertDTO.UserID.String())
	}

	logging.LogSuccess(KeyCreateSubscription, MsgSubscriptionCreated, map[string]interface{}{
		"subscription_id": subscription.ID,
	})

	return response.Created(c, MsgSubscriptionCreated, subscription)
}

// @Summary Get My Subscription
// @Description Retrieves the subscription details for the authenticated user.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.ApiResponse "Subscription retrieved successfully"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Failure 500 {object} response.ApiResponse "Internal server error"
// @Router /v1/api/subscriptions/me [get]
func (sc *SubscriptionController) GetMySubscription(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return response.Unauthorized(c, err.Error(), MsgUnauthorized)
	}

	logging.LogIncomingRequest(c, KeyGetMySubscription)

	subscription, err := sc.service.GetSubscriptionByUser(context.Background(), userID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyGetMySubscription, userID.String())
	}

	logging.LogSuccess(KeyGetMySubscription, MsgSubscriptionRetrieved, map[string]interface{}{
		"subscription_id": subscription.ID,
	})

	return response.OK(c, MsgSubscriptionRetrieved, subscription)
}

// @Summary Cancel My Subscription
// @Description Cancels the subscription for the authenticated user.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param lesson_id path string true "Lesson ID"
// @Success 200 {object} response.ApiResponse "Subscription cancelled successfully"
// @Failure 400 {object} response.ApiResponse "Invalid parameter"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Failure 500 {object} response.ApiResponse "Internal server error"
// @Router /v1/api/subscriptions/cancel/{lesson_id} [post]
func (sc *SubscriptionController) CancelMySubscription(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return response.Unauthorized(c, err.Error(), MsgUnauthorized)
	}

	logging.LogIncomingRequest(c, KeyCancelMySubscription)

	subscriptionID, err := response.GetUUIDParam(c, "lesson_id", KeyCancelMySubscription)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidParam)
	}

	err = sc.service.CancelSubscription(context.Background(), userID, subscriptionID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyCancelMySubscription, userID.String())
	}

	logging.LogSuccess(KeyCancelMySubscription, MsgSubscriptionCancelled, map[string]interface{}{
		"subscription_id": subscriptionID,
	})

	return response.OK(c, MsgSubscriptionCancelled, nil)
}
