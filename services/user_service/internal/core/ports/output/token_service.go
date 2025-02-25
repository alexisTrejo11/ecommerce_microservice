package output

import (
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/tokens"
)

type TokenService interface {
	GenerateTokens(userID, email, role string) (string, string, error)
	VerifyToken(tokenString string, tokenType tokens.TokenType) (*tokens.Claims, error)
	RefreshToken(refreshToken string) (string, error)
	GetTokenExpirationDate(tokenString string, tokenType tokens.TokenType) (time.Time, error)
}
