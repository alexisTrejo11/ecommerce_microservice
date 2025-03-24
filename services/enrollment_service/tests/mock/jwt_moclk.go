package mocks

import (
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	tokens "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/token"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockJWTManager struct {
	jwt.JWTManager
	mock.Mock
}

func (m *MockJWTManager) ExtractAndValidateToken(c *fiber.Ctx) (*tokens.Claims, error) {
	// Implement mock behavior
	claims := &tokens.Claims{
		UserID: uuid.MustParse("ff8c6bc9-2a2c-4ab9-b65f-88deb761b5de").String(),
	}
	return claims, nil
}

func (m *MockJWTManager) GetUserIDFromToken(c *fiber.Ctx) (uuid.UUID, error) {
	// Implement mock behavior
	return uuid.MustParse("ff8c6bc9-2a2c-4ab9-b65f-88deb761b5de"), nil
}

func (m *MockJWTManager) VerifyToken(tokenString string) (*tokens.Claims, error) {
	// Implement mock behavior
	claims := &tokens.Claims{
		UserID: uuid.MustParse("ff8c6bc9-2a2c-4ab9-b65f-88deb761b5de").String(),
	}
	return claims, nil
}

func (m *MockJWTManager) GetTokenExpirationDate(tokenString string) (time.Time, error) {
	// Implement mock behavior
	return time.Now().Add(24 * time.Hour), nil
}
