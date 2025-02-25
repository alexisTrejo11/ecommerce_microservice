package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/config"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/tokens"
)

type TokenServiceImpl struct {
	jwtManager   *jwt.JWTManager
	tokenFactory tokens.TokenFactory
}

func NewTokenService(jwtManager *jwt.JWTManager) output.TokenService {
	return &TokenServiceImpl{
		jwtManager: jwtManager,
	}
}

func (s *TokenServiceImpl) GenerateTokens(userID, email, role string) (string, string, error) {
	accesTokenFactory, _ := s.tokenFactory.CreateToken(tokens.AccessTokenENUM)
	accessToken, err := accesTokenFactory.Generate(userID, email, role)
	if err != nil {
		return "", "", err
	}

	refreshTokenFactory, _ := s.tokenFactory.CreateToken(tokens.RefreshTokenENUM)
	refreshToken, err := refreshTokenFactory.Generate(userID, email, role)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *TokenServiceImpl) VerifyToken(tokenString string, tokenType tokens.TokenType) (*tokens.Claims, error) {
	factory, err := s.tokenFactory.CreateToken(tokenType)
	if err != nil {
		return nil, err
	}

	claims, err := factory.VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *TokenServiceImpl) RefreshToken(refreshToken string) (string, error) {
	claims, err := s.VerifyToken(refreshToken, tokens.AccessTokenENUM)
	if err != nil {
		return "", err
	}

	accessToken, err := s.jwtManager.GenerateToken(claims.UserID, claims.Email, claims.Role)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *TokenServiceImpl) GetTokenExpirationDate(tokenString string, tokenType tokens.TokenType) (time.Time, error) {
	factory, err := s.tokenFactory.CreateToken(tokenType)
	if err != nil {
		return time.Time{}, err
	}

	expirationDate, err := factory.GetTokenExpirationDate(tokenString)
	if err != nil {
		return time.Time{}, err
	}
	return expirationDate, nil
}

func (s *TokenServiceImpl) BlacklistToken(token string, expiresIn time.Duration) error {
	ctx := context.Background()
	err := config.RedisClient.Set(ctx, "blacklist:"+token, true, expiresIn).Err()
	if err != nil {
		return fmt.Errorf("error al agregar el token a la lista negra: %v", err)
	}
	return nil
}

func (s *TokenServiceImpl) IsTokenBlacklisted(token string) bool {
	ctx := context.Background()
	exists, err := config.RedisClient.Exists(ctx, "blacklist:"+token).Result()
	if err != nil {
		panic("error al verificar la lista negra: " + err.Error())
	}
	return exists == 1
}
