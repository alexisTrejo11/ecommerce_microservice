package response

import (
	"errors"

	appErr "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/error"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func HandleApplicationError(c *fiber.Ctx, err error, action string, resourceID string) error {
	var application *appErr.ApplicationError
	if errors.As(err, &application) {
		logError(action, resourceID, application.Code, application.Message)

		switch application.Code {
		case "CERTIFCATE_NOT_FOUND", "ENROLLMENT_NOT_FOUND", "SUBSCRIPTION_NOT_FOUND", "PROGRESS_NOT_FOUND":
			return Error(c, fiber.StatusNotFound, application.Message, application.Code)
		case "DATABASE_ERROR":
			return Error(c, fiber.StatusInternalServerError, application.Message, application.Code)
		default:
			return Error(c, fiber.StatusInternalServerError, "An unexpected error occurred", application.Code)
		}
	}

	logError(action, resourceID, "UNKNOWN_ERROR", err.Error())
	return Error(c, fiber.StatusInternalServerError, err.Error(), "UNKNOWN_ERROR")
}

func logError(action string, resourceID string, code string, message string) {
	log.Error().
		Str("action", action).
		Str("status", "failed").
		Str("resource_id", resourceID).
		Str("code", code).
		Str("error", message).
		Msg("Application error occurred")
}
