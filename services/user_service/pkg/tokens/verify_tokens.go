package tokens

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/config"
	"github.com/go-redis/redis/v8"
)

type VerifyToken struct{}

func (t *VerifyToken) Generate(email, userId, role string) (string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	tokenInt := rand.Intn(899999999) + 100000000
	token := strconv.Itoa(tokenInt)

	ctx := context.Background()
	claims := Claims{
		Email:     email,
		UserID:    userId,
		Role:      role,
		TokenType: "VERIFY_TOKEN",
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}

	jsonData, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("error marshaling token data: %w", err)
	}

	key := fmt.Sprintf("verification_token:%s", token)

	if err := config.RedisClient.Set(ctx, key, jsonData, time.Until(claims.ExpiresAt)).Err(); err != nil {
		return "", fmt.Errorf("error storing verification token: %w", err)
	}

	return token, nil
}

func (t *VerifyToken) VerifyToken(tokenString string) (*Claims, error) {
	ctx := context.Background()
	key := fmt.Sprintf("verification_token:%s", tokenString)

	jsonData, err := config.RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("token not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error retrieving token: %w", err)
	}

	var claims Claims
	if err := json.Unmarshal([]byte(jsonData), &claims); err != nil {
		return nil, fmt.Errorf("error unmarshaling claims: %w", err)
	}

	if time.Now().After(claims.ExpiresAt) {
		config.RedisClient.Del(ctx, key)
		return nil, fmt.Errorf("token expired")
	}

	config.RedisClient.Del(ctx, key)
	return &claims, nil
}

func (t *VerifyToken) GetTokenExpirationDate(tokenString string) (time.Time, error) {
	claims, err := t.VerifyToken(tokenString)
	if err != nil {
		return time.Time{}, err
	}
	return claims.ExpiresAt, nil
}
