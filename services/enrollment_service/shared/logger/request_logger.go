package logging

import (
	"github.com/gofiber/fiber/v2"
)

func LogIncomingRequest(c *fiber.Ctx, action string) {
	Logger.Info().
		Str("action", action).
		Str("method", c.Method()).
		Str("route", c.Route().Path).
		Str("ip", c.IP()).
		Interface("user", c.Locals("user_id")).
		Msg("Incoming request")
}

func LogIncomingRequestWithPayload(c *fiber.Ctx, action string) {
	Logger.Info().
		Str("action", action).
		Str("method", c.Method()).
		Str("route", c.Route().Path).
		Str("ip", c.IP()).
		Interface("user", c.Locals("user_id")).
		RawJSON("payload", c.Body()).
		Msg("Incoming request")
}

func LogError(action, message string, fields map[string]interface{}) {
	event := Logger.Error().
		Str("action", action).
		Str("status", "failed")

	for k, v := range fields {
		event = event.Interface(k, v)
	}

	event.Msg(message)
}

func LogSuccess(action, message string, fields map[string]interface{}) {
	event := Logger.Info().
		Str("action", action).
		Str("status", "success")

	for k, v := range fields {
		event = event.Interface(k, v)
	}

	event.Msg(message)
}
