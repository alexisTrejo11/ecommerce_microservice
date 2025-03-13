package response

import (
	"errors"

	customErrors "github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/errors"
	logging "github.com/alexisTrejo11/ecommerce_microservice/course-service/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func HandleApplicationError(c *fiber.Ctx, err error, action string, resourceID string) error {
	var domainErr *customErrors.DomainError
	if errors.As(err, &domainErr) {
		logError(action, resourceID, domainErr.Code, domainErr.Message)

		switch domainErr.Code {
		case "COURSE_NOT_FOUND", "LESSON_NOT_FOUND", "MODULE_NOT_FOUND", "RESOURCE_NOT_FOUND":
			return Error(c, fiber.StatusNotFound, domainErr.Message, domainErr.Code)
		case "DATABASE_ERROR":
			return Error(c, fiber.StatusInternalServerError, domainErr.Message, domainErr.Code)
		default:
			return Error(c, fiber.StatusInternalServerError, "An unexpected error occurred", domainErr.Code)
		}
	}

	logError(action, resourceID, "UNKNOWN_ERROR", err.Error())
	return Error(c, fiber.StatusInternalServerError, "An unknown error occurred", "UNKNOWN_ERROR")
}

func logError(action string, resourceID string, code string, message string) {
	logging.Logger.WithFields(logrus.Fields{
		"action":      action,
		"status":      "failed",
		"resource_id": resourceID,
		"code":        code,
		"error":       message,
	}).Error("Application error occurred")
}
