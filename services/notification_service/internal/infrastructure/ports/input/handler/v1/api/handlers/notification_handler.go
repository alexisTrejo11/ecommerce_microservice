package handler

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/response"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/utils"
	logging "github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/log"
	"github.com/gofiber/fiber/v2"
)

type NotificationHandler struct {
	notificationUseCase input.NotificationUseCase
}

func NewNotificationHandler(notificationUseCase input.NotificationUseCase) *NotificationHandler {
	return &NotificationHandler{
		notificationUseCase: notificationUseCase,
	}
}

func (h *NotificationHandler) GetNotificationByUserId(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_notification_by_user_id")

	userId, err := utils.GetUUIDParam(c, "user_id")
	if err != nil {
		logging.LogError("get_notification_by_user_id", "Invalid user ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	pageable, err := utils.GetPageData(c)
	if err != nil {
		logging.LogError("get_notification_by_user_id", "Invalid page data", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid_page_data")
	}

	notifications, _, err := h.notificationUseCase.GetUserNotifications(context.Background(), userId, *pageable)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_notification_by_user_id", userId.String())
	}

	logging.LogSuccess("get_notification_by_user_id", "User Notification Successfully Retrieved", map[string]interface{}{
		"user_id": userId,
	})

	return response.OK(c, "User Notification Successfully Retrieved", notifications)
}

func (h *NotificationHandler) GetNotificationById(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_notification_by_id")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("get_notification_by_id", "invalid notification ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid notification ID")
	}

	notification, err := h.notificationUseCase.GetNotification(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_notification_by_id", id.String())
	}

	logging.LogSuccess("get_notification_by_id", "Notification Successfully Retrieved", map[string]interface{}{
		"notification_id": id,
	})

	return response.OK(c, "Notification Successfully Retrieved", notification)
}

func (h *NotificationHandler) CancelNotification(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "cancel_notification")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError(" cancel_notifaction", "invalid notification ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid notification ID")
	}

	err = h.notificationUseCase.CancelNotification(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "cancel_notification", id.String())
	}

	logging.LogSuccess(" cancel_course", "Notification Successfully Cancelled", map[string]interface{}{
		"notification_id": id,
	})

	return response.OK(c, "Notification Successfully Cancelled", nil)
}
