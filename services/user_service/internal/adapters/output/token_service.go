package repository

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

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

func (s *TokenService) GetTokenExpirationDate(tokenString string) (time.Time, error) {
	expirationDate, err := s.jwtManager.GetTokenExpirationDate(tokenString)
	if err != nil {
		return time.Time{}, err
	}
	return expirationDate, nil
}

func GenerateActivationToken() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	min := 100000000
	max := 999999999
	token := rand.Intn(max-min) + min
	return fmt.Sprintf("%09d", token)
}

type ActivationTokenStore struct {
	tokens map[string]TokenData
	mu     sync.Mutex
}

type TokenData struct {
	email     string
	ExpiresAt time.Time
}

func NewTokenStore() *ActivationTokenStore {
	return &ActivationTokenStore{
		tokens: make(map[string]TokenData),
	}
}

func (ts *ActivationTokenStore) SaveActivationToken(token string, email string, expiresAt time.Time) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.tokens[token] = TokenData{
		email:     email,
		ExpiresAt: expiresAt,
	}
}

func (ts *ActivationTokenStore) ValidarToken(token string) (string, bool) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	data, exists := ts.tokens[token]
	if !exists {
		return "", false
	}

	if time.Now().After(data.ExpiresAt) {
		delete(ts.tokens, token)
		return "", false
	}

	return data.email, true
}
