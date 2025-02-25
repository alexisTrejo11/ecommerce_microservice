package response

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	Code      int         `json:"code"`
}

func Success(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(ApiResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
		Code:      statusCode,
	})
}

func Error(c *fiber.Ctx, statusCode int, message string, err interface{}) error {
	return c.Status(statusCode).JSON(ApiResponse{
		Success:   false,
		Message:   message,
		Errors:    err,
		Timestamp: time.Now(),
		Code:      statusCode,
	})
}

func OK(c *fiber.Ctx, message string, data interface{}) error {
	return Success(c, http.StatusOK, message, data)
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return Success(c, http.StatusCreated, message, data)
}

func BadRequest(c *fiber.Ctx, message string, err interface{}) error {
	return Error(c, http.StatusBadRequest, message, err)
}

func Unauthorized(c *fiber.Ctx, message string, err interface{}) error {
	return Error(c, http.StatusUnauthorized, message, err)
}

func Forbidden(c *fiber.Ctx, message string, err interface{}) error {
	return Error(c, http.StatusForbidden, message, err)
}

func NotFound(c *fiber.Ctx, message string, err interface{}) error {
	return Error(c, http.StatusNotFound, message, err)
}

func InternalServerError(c *fiber.Ctx, message string, err interface{}) error {
	return Error(c, http.StatusInternalServerError, message, err)
}
