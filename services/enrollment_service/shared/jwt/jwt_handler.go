package jwt

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	tokens "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/token"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager interface {
	VerifyToken(tokenString string) (*tokens.Claims, error)
	ExtractAndValidateToken(c *fiber.Ctx) (*tokens.Claims, error)
	GetTokenExpirationDate(tokenString string) (time.Time, error)
	GetUserIDFromToken(c *fiber.Ctx) (uuid.UUID, error)
}

type JWTManager struct {
	config tokens.Config
}

func NewJWTManager() (TokenManager, error) {
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

func (j *JWTManager) GetUserIDFromToken(c *fiber.Ctx) (uuid.UUID, error) {
	claims, err := j.ExtractAndValidateToken(c)
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return uuid.Nil, err
	}

	c.Request().Header.Set("user", userID.String())

	return userID, nil
}
