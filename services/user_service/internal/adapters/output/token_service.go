package repository

import (
	jwt "github.com/alexisTrejo11/ecommerce_microservice/pkg/jwt"
)

type TokenService struct {
	jwtManager *jwt.JWTManager
}

func NewTokenService(jwtManager *jwt.JWTManager) *TokenService {
	return &TokenService{
		jwtManager: jwtManager,
	}
}

func (s *TokenService) GenerateTokens(userID, email, role string) (string, string, error) {
	accessToken, err := s.jwtManager.GenerateToken(userID, email, role)
	if err != nil {
		return "", "", err
	}

	// Generamos un refresh token sin incluir email ni role
	refreshToken, err := s.jwtManager.GenerateToken(userID, "", "")
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *TokenService) VerifyToken(tokenString string) (*jwt.Claims, error) {
	return s.jwtManager.VerifyToken(tokenString)
}

func (s *TokenService) RefreshToken(refreshToken string) (string, error) {
	claims, err := s.VerifyToken(refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, err := s.jwtManager.GenerateToken(claims.UserID, claims.Email, claims.Role)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
