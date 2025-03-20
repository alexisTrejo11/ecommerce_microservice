package middleware

import (
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const UserIDKey = "userId"

func JWTAuthMiddleware(jwtManager jwt.JWTManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := jwtManager.GetUserIDFromToken(c)
		if err != nil {
			logging.LogError("jwt_authentication", "Invalid JWT token", map[string]interface{}{
				"error": err.Error(),
			})
			return response.BadRequest(c, err.Error(), "invalid_user_id")
		}

		c.Locals(UserIDKey, userID)

		log.Info().
			Str("action", "jwt_authentication").
			Str("status", "success").
			Str("user_id", userID.String()).
			Msg("User authenticated successfully")

		return c.Next()
	}
}

func GetUserID(c *fiber.Ctx) (uuid.UUID, error) {
	userID, ok := c.Locals(UserIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}
	return userID, nil
}
