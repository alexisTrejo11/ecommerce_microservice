package output

import "github.com/alexisTrejo11/ecommerce_microservice/pkg/jwt"

type TokenService interface {
	GenerateTokens(userID, email, role string) (string, string, error)
	VerifyToken(tokenString string) (*jwt.Claims, error)
	RefreshToken(refreshToken string) (string, error)
}
