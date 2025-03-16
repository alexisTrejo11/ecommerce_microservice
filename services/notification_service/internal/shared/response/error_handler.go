package response

import (
	"errors"
	"strings"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	repository "github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/ports/output"
	logging "github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/log"
	"github.com/gofiber/fiber/v2"
)

func HandleApplicationError(c *fiber.Ctx, err error, action string, resourceID string) error {
	var notificationErr *domain.NotificationError
	var repositoryErr *repository.RepositoryError

	if errors.As(err, &notificationErr) || errors.As(err, &repositoryErr) {
		logError(action, resourceID, notificationErr.Code, notificationErr.Message)

		switch {
		case notificationErr.Code == "NOTIFICATION_NOT_FOUND" || notificationErr.Code == "USER_NOT_FOUND":
			return Error(c, fiber.StatusNotFound, notificationErr.Message, notificationErr.Code)
		case notificationErr.Code == "DATABASE_ERROR":
			return Error(c, fiber.StatusInternalServerError, notificationErr.Message, notificationErr.Code)
		case strings.Contains(notificationErr.Code, "NOTIFICATION_INVALID"):
			return Error(c, fiber.StatusBadRequest, notificationErr.Message, notificationErr.Code)
		default:
			return Error(c, fiber.StatusInternalServerError, "An unexpected error occurred", notificationErr.Code)
		}
	}

	logError(action, resourceID, "UNKNOWN_ERROR", err.Error())
	return Error(c, fiber.StatusInternalServerError, "An unknown error occurred", "UNKNOWN_ERROR")
}

func logError(action string, resourceID string, code string, message string) {
	logging.Logger.Error().
		Str("action", action).
		Str("status", "failed").
		Str("resource_id", resourceID).
		Str("code", code).
		Str("error", message).
		Msg("Application error occurred")
}
