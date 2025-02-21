package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/config"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	jwt "github.com/alexisTrejo11/ecommerce_microservice/pkg/jwt"
	"github.com/go-redis/redis/v8"
)

type TokenServiceImpl struct {
	jwtManager *jwt.JWTManager
}

func NewTokenService(jwtManager *jwt.JWTManager) output.TokenService {
	return &TokenServiceImpl{
		jwtManager: jwtManager,
	}
}

func (s *TokenServiceImpl) GenerateTokens(userID, email, role string) (string, string, error) {
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

func (s *TokenServiceImpl) VerifyToken(tokenString string) (*jwt.Claims, error) {
	return s.jwtManager.VerifyToken(tokenString)
}

func (s *TokenServiceImpl) RefreshToken(refreshToken string) (string, error) {
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

func (s *TokenServiceImpl) GetTokenExpirationDate(tokenString string) (time.Time, error) {
	expirationDate, err := s.jwtManager.GetTokenExpirationDate(tokenString)
	if err != nil {
		return time.Time{}, err
	}
	return expirationDate, nil
}

func (s *TokenServiceImpl) generatectivationToken() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	min := 100000000
	max := 999999999
	token := rand.Intn(max-min) + min
	return fmt.Sprintf("%09d", token)
}

type TokenData struct {
	Email     string    `json:"email"`
	UserID    string    `json:"user_id"`
	Role      string    `json:"role"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (s *TokenServiceImpl) SaveActivationToken(token, email string, expiresAt time.Time) error {
	fmt.Println("Hola1")
	ctx := context.Background()

	tokenData := TokenData{
		Email:     email,
		ExpiresAt: expiresAt,
	}

	fmt.Println("Hola2")

	jsonData, err := json.Marshal(tokenData)
	if err != nil {
		return fmt.Errorf("error marshaling token data: %v", err)
	}

	fmt.Println("Hola3")
	key := fmt.Sprintf("activation_token:%s", token)

	err = config.RedisClient.Set(ctx, key, jsonData, time.Until(expiresAt)).Err()
	if err != nil {
		return fmt.Errorf("error storing activation token: %v", err)
	}
	fmt.Print("Hola4")

	return nil
}

func (s *TokenServiceImpl) ValidateActivationToken(token string) (string, bool) {
	ctx := context.Background()

	key := fmt.Sprintf("activation_token:%s", token)
	jsonData, err := config.RedisClient.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", false
	}

	if err != nil {
		return "", false
	}

	var tokenData TokenData
	err = json.Unmarshal([]byte(jsonData), &tokenData)
	if err != nil {
		return "", false
	}

	if time.Now().After(tokenData.ExpiresAt) {
		config.RedisClient.Del(ctx, key)
		return "", false
	}

	config.RedisClient.Del(ctx, key)

	return tokenData.Email, true
}

func (s *TokenServiceImpl) GetActivationToken(userID, email, role string) string {
	activationToken := s.generatectivationToken()

	err := s.SaveActivationToken(activationToken, email, time.Now().Add(time.Hour*24))
	if err != nil {
		fmt.Printf("Error saving token: %v\n", err)
		return ""
	}
	return activationToken
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
