package middleware

import (
	"context"
	"time"

	port "github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func LoggerMiddleware(log port.LoggerPort) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Generar o extraer requestID
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
			c.Set("X-Request-ID", requestID)
		}

		ctx := context.WithValue(c.Context(), "requestID", requestID)
		c.SetUserContext(ctx)

		path := c.Path()
		method := c.Method()
		ip := c.IP()
		userAgent := c.Get("User-Agent")

		log.Info(ctx, "Request started",
			port.Field{Key: "method", Value: method},
			port.Field{Key: "path", Value: path},
			port.Field{Key: "ip", Value: ip},
			port.Field{Key: "user_agent", Value: userAgent},
		)

		err := c.Next()

		latency := time.Since(start).Milliseconds()
		status := c.Response().StatusCode()

		logFields := []port.Field{
			{Key: "method", Value: method},
			{Key: "path", Value: path},
			{Key: "status", Value: status},
			{Key: "latency_ms", Value: latency},
			{Key: "ip", Value: ip},
		}

		if err != nil {
			logFields = append(logFields, port.Field{Key: "error", Value: err.Error()})
		}

		if status >= 500 {
			log.Error(ctx, "Request failed", logFields...)
		} else if status >= 400 {
			log.Warn(ctx, "Request completed with client error", logFields...)
		} else {
			log.Info(ctx, "Request completed successfully", logFields...)
		}

		return err
	}
}

func RequestIDMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
			c.Set("X-Request-ID", requestID)
		}
		ctx := context.WithValue(c.Context(), "requestID", requestID)
		c.SetUserContext(ctx)

		return c.Next()
	}
}
