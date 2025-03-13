package logging

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func LogIncomingRequest(c *fiber.Ctx, action string) {
	Logger.WithFields(logrus.Fields{
		"action": action,
		"method": c.Method(),
		"route":  c.Route().Path,
		"ip":     c.IP(),
		"user":   c.Locals("user_id"),
	}).Info("Incoming request")
}

func LogIncomingRequestWithPayload(c *fiber.Ctx, action string) {
	Logger.WithFields(logrus.Fields{
		"action":  action,
		"method":  c.Method(),
		"route":   c.Route().Path,
		"ip":      c.IP(),
		"user":    c.Locals("user_id"),
		"payload": c.Body(),
	}).Info("Incoming request")
}

func LogError(action, message string, fields map[string]interface{}) {
	entry := Logger.WithFields(logrus.Fields{
		"action": action,
		"status": "failed",
	})

	for k, v := range fields {
		entry = entry.WithField(k, v)
	}

	entry.Error(message)
}

func LogSuccess(action, message string, fields map[string]interface{}) {
	entry := Logger.WithFields(logrus.Fields{
		"action": action,
		"status": "success",
	})

	for k, v := range fields {
		entry = entry.WithField(k, v)
	}

	entry.Info(message)
}
