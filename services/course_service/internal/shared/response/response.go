package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// ApiResponse represents the structure of the API response.
// It contains the success status, message, data, errors (if any), timestamp, and HTTP status code.
// It supports pagination and correlation tracking.
type ApiResponse[T any, E any] struct {
	Success       bool        `json:"success"`
	Message       interface{} `json:"message,omitempty"` // Puede ser un string o lista de mensajes
	Data          T           `json:"data,omitempty"`
	Errors        E           `json:"errors,omitempty"`
	Timestamp     time.Time   `json:"timestamp"`
	Code          int         `json:"code"`
	CorrelationID string      `json:"correlationId,omitempty"` // Para seguimiento de solicitudes
	TotalCount    int         `json:"totalCount,omitempty"`    // Para paginación
	Page          int         `json:"page,omitempty"`          // Página actual
	PerPage       int         `json:"perPage,omitempty"`       // Elementos por página
}

// Success sends a success response with the given data.
// @Summary Send a success response
// @Description Send a success response with a custom message and data
// @Param message query string true "Response message"
// @Param data query object false "Response data"
// @Success 200 {object} ApiResponse "Success response"
func Success[T any](c *fiber.Ctx, statusCode int, message interface{}, data T) error {
	return c.Status(statusCode).JSON(ApiResponse[T, any]{
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
func Error[E any](c *fiber.Ctx, statusCode int, message string, err E) error {
	return c.Status(statusCode).JSON(ApiResponse[any, E]{
		Success:   false,
		Message:   message,
		Errors:    err,
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
func OK[T any](c *fiber.Ctx, message string, data T) error {
	return Success[T](c, 200, message, data)
}

// Created sends a response indicating that a resource was created.
// @Summary Send a created response
// @Description Send a response indicating a resource was created successfully
// @Param message query string true "Response message"
// @Param data query object false "Response data"
// @Success 201 {object} ApiResponse "Created response"
func Created[T any](c *fiber.Ctx, message string, data T) error {
	return Success[T](c, 201, message, data)
}

// BadRequest sends a response indicating that the request was invalid.
// @Summary Send a bad request response
// @Description Send a bad request response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 400 {object} ApiResponse "Bad request response"
func BadRequest[E any](c *fiber.Ctx, message string, err E) error {
	return Error[E](c, 400, message, err)
}

// Unauthorized sends a response indicating that the user is not authorized.
// @Summary Send an unauthorized response
// @Description Send an unauthorized response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 401 {object} ApiResponse "Unauthorized response"
func Unauthorized[E any](c *fiber.Ctx, message string, err E) error {
	return Error[E](c, 401, message, err)
}

// Forbidden sends a response indicating that access is forbidden.
// @Summary Send a forbidden response
// @Description Send a forbidden response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 403 {object} ApiResponse "Forbidden response"
func Forbidden[E any](c *fiber.Ctx, message string, err E) error {
	return Error[E](c, 403, message, err)
}

// NotFound sends a response indicating that the resource was not found.
// @Summary Send a not found response
// @Description Send a not found response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 404 {object} ApiResponse "Not found response"
func NotFound[E any](c *fiber.Ctx, message string, err E) error {
	return Error[E](c, 404, message, err)
}

// InternalServerError sends a response indicating an internal server error.
// @Summary Send an internal server error response
// @Description Send an internal server error response with an error message
// @Param message query string true "Response message"
// @Param errors query object true "Error details"
// @Failure 500 {object} ApiResponse "Internal server error response"
func InternalServerError[E any](c *fiber.Ctx, message string, err E) error {
	return Error[E](c, 500, message, err)
}
