package tokens

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessToken struct {
}

func (t *AccessToken) Generate(email, userId, role string) (string, error) {
	config := getConfig()
	claims := Claims{
		UserID:    userId,
		Email:     email,
		Role:      role,
		TokenType: "ACCESS_TOKEN",
		ExpiresAt: time.Now().Add(config.RefreshTokenExpiry),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

func (t *AccessToken) VerifyToken(tokenString string) (*Claims, error) {
	return verifyJWTToken(tokenString)
}

func (t *AccessToken) GetTokenExpirationDate(tokenString string) (time.Time, error) {
	return getJWTTokenExpirationDate(tokenString)
}

type RefreshTokenFactory struct{}

func (f *RefreshTokenFactory) CreateToken() Token {
	return &RefreshToken{}
}

type RefreshToken struct{}

func (t *RefreshToken) Generate(email, userId, role string) (string, error) {
	config := getConfig()
	claims := Claims{
		UserID:    userId,
		Email:     email,
		Role:      role,
		TokenType: "REFRESH_TOKEN",
		ExpiresAt: time.Now().Add(config.RefreshTokenExpiry),
	}

	fmt.Println("Claims", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

func (t *RefreshToken) VerifyToken(tokenString string) (*Claims, error) {
	return verifyJWTToken(tokenString)
}

func (t *RefreshToken) GetTokenExpirationDate(tokenString string) (time.Time, error) {
	return getJWTTokenExpirationDate(tokenString)
}

func verifyJWTToken(tokenString string) (*Claims, error) {
	config := getConfig()
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func getJWTTokenExpirationDate(tokenString string) (time.Time, error) {
	claims, err := verifyJWTToken(tokenString)
	if err != nil {
		return time.Time{}, err
	}
	return claims.ExpiresAt, nil
}
