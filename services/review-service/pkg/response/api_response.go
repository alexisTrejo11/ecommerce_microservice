package response

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ApiResponse represents the structure of the API response.
// It contains the success status, message, data, errors (if any), timestamp, and HTTP status code.
// It supports pagination and correlation tracking.
type ApiResponse struct {
	Success       bool        `json:"success"`
	Message       interface{} `json:"message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	ErrorCode     interface{} `json:"error_code,omitempty"`
	Timestamp     time.Time   `json:"timestamp"`
	Code          int         `json:"code"`
	CorrelationID string      `json:"correlationId,omitempty"`
	TotalCount    int         `json:"totalCount,omitempty"`
	Page          int         `json:"page,omitempty"`
	PerPage       int         `json:"perPage,omitempty"`
}

// Success sends a success response with the given data.
// @Summary Send a success response
// @Description Send a success response with a custom message and data
// @Param message query string true "Response message"
// @Param data query object false "Response data"
// @Success 200 {object} ApiResponse "Success response"
func Success(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(ApiResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
		Code:      statusCode,
	})
}

// Error sends an error response with the given error details.
// @Summary Send an error response
// @Description Send an error response with a custom message and error details
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 400 {object} ApiResponse "Error response"
func Error(c *fiber.Ctx, statusCode int, message interface{}, errCode interface{}) error {
	return c.Status(statusCode).JSON(ApiResponse{
		Success:   false,
		Message:   message,
		ErrorCode: errCode,
		Timestamp: time.Now(),
		Code:      statusCode,
	})
}

// OK sends an OK response with the given message and data.
// @Summary Send an OK response
// @Description Send an OK response with a message and data
// @Param message query string true "Response message"
// @Param data query object false "Response data"
// @Success 200 {object} ApiResponse "OK response"
func OK(c *fiber.Ctx, message string, data interface{}) error {
	return Success(c, http.StatusOK, message, data)
}

// Created sends a response indicating that a resource was created.
// @Summary Send a created response
// @Description Send a response indicating a resource was created successfully
// @Param message query string true "Response message"
// @Param data query object false "Response data"
// @Success 201 {object} ApiResponse "Created response"
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return Success(c, http.StatusCreated, message, data)
}

// BadRequest sends a response indicating that the request was invalid.
// @Summary Send a bad request response
// @Description Send a bad request response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 400 {object} ApiResponse "Bad request response"
func BadRequest(c *fiber.Ctx, message interface{}, errCode interface{}) error {
	return Error(c, http.StatusBadRequest, message, errCode)
}

// Unauthorized sends a response indicating that the user is not authorized.
// @Summary Send an unauthorized response
// @Description Send an unauthorized response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 401 {object} ApiResponse "Unauthorized response"
func Unauthorized(c *fiber.Ctx, message string, errCode interface{}) error {
	return Error(c, http.StatusUnauthorized, message, errCode)
}

// Forbidden sends a response indicating that access is forbidden.
// @Summary Send a forbidden response
// @Description Send a forbidden response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 403 {object} ApiResponse "Forbidden response"
func Forbidden(c *fiber.Ctx, message string, errCode interface{}) error {
	return Error(c, http.StatusForbidden, message, errCode)
}

// NotFound sends a response indicating that the resource was not found.
// @Summary Send a not found response
// @Description Send a not found response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 404 {object} ApiResponse "Not found response"
func NotFound(c *fiber.Ctx, message string, errCode interface{}) error {
	return Error(c, http.StatusNotFound, message, errCode)
}

// InternalServerError sends a response indicating an internal server error.
// @Summary Send an internal server error response
// @Description Send an internal server error response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 500 {object} ApiResponse "Internal server error response"
func InternalServerError(c *fiber.Ctx, message string, errCode interface{}) error {
	return Error(c, http.StatusInternalServerError, message, errCode)
}
