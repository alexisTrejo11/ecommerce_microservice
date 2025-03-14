package handler

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/response"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/utils"
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
	userId, err := utils.GetUUIDParam(c, "user_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	pageable, err := utils.GetPageData(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_page_data")
	}

	notifications, _, err := h.notificationUseCase.GetUserNotifications(context.Background(), userId, *pageable)
	if err != nil {
		return response.Error(c, 400, err.Error(), "invalid_page_data")
	}

	return response.OK(c, "User Notification Successfully Fetched", notifications)
}

func (h *NotificationHandler) GetNotificationById(c *fiber.Ctx) error {
	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_notification_id")
	}

	notification, err := h.notificationUseCase.GetNotification(context.Background(), id)
	if err != nil {
		return response.Error(c, 404, err.Error(), "notification_not_found")
	}

	return response.OK(c, "Notification Successfully Retrieved", notification)
}

func (h *NotificationHandler) DeleteNotification(c *fiber.Ctx) error {
	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_notification_id")
	}

	err = h.notificationUseCase.CancelNotification(context.Background(), id)
	if err != nil {
		return response.Error(c, 404, err.Error(), "notification_not_found")
	}

	return response.OK(c, "Notification Successfully Delete", nil)
}
