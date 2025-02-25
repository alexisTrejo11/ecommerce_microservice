package jwt

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/tokens"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	config tokens.Config
}

func NewJWTManager() (*JWTManager, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return nil, errors.New("JWT_SECRET_KEY is not defined in the environment variables")
	}

	accessTokenExpiry, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRY"))
	if err != nil {
		accessTokenExpiry = time.Minute * 15
	}

	refreshTokenExpiry, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRY"))
	if err != nil {
		refreshTokenExpiry = time.Hour * 24 * 7
	}

	config := tokens.Config{
		AccessTokenExpiry:  accessTokenExpiry,
		RefreshTokenExpiry: refreshTokenExpiry,
		JWTSecret:          []byte(secret),
	}

	return &JWTManager{config: config}, nil
}

func (j *JWTManager) GenerateToken(userID, email, role string) (string, error) {
	claims := tokens.Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.AccessTokenExpiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.config.JWTSecret)
}

func (j *JWTManager) VerifyToken(tokenString string) (*tokens.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokens.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.config.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokens.Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func (j *JWTManager) ExtractAndValidateToken(c *fiber.Ctx) (*tokens.Claims, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("invalid authorization header format")
	}

	tokenString := parts[1]

	claims, err := j.VerifyToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}

func (j *JWTManager) GetTokenExpirationDate(tokenString string) (time.Time, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokens.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.config.JWTSecret, nil
	})
	if err != nil {
		return time.Time{}, err
	}

	claims, ok := token.Claims.(*tokens.Claims)
	if !ok || !token.Valid {
		return time.Time{}, jwt.ErrSignatureInvalid
	}

	return claims.ExpiresAt, nil
}
