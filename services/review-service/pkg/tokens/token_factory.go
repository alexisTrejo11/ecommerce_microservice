package tokens

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessTokenENUM  TokenType = "ACCESS_TOKEN"
	RefreshTokenENUM TokenType = "REFRESH_TOKEN"
	VerifyTokenENUM  TokenType = "VERIFY_TOKEN"
)

type Token interface {
	Generate(email, userId, role string) (string, error)
	VerifyToken(tokenString string) (*Claims, error)
	GetTokenExpirationDate(tokenString string) (time.Time, error)
}

type TokenFactory struct{}

func NewTokenFactory() *TokenFactory {
	return &TokenFactory{}
}

func (f *TokenFactory) CreateToken(tokenType TokenType) (Token, error) {
	switch tokenType {
	case AccessTokenENUM:
		return &AccessToken{}, nil
	case RefreshTokenENUM:
		return &RefreshToken{}, nil
	default:
		return nil, fmt.Errorf("unsupported token type: %s", tokenType)
	}
}

type Claims struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	TokenType string    `json:"token_type"`
	ExpiresAt time.Time `json:"expires_at"`
	jwt.RegisteredClaims
}

type Config struct {
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
	JWTSecret          []byte
}

func getConfig() Config {
	secret := os.Getenv("JWT_SECRET_KEY")
	accessTokenExpiry, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRY"))
	if err != nil {
		accessTokenExpiry = time.Minute * 15
	}

	refreshTokenExpiry, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRY"))
	if err != nil {
		refreshTokenExpiry = time.Hour * 24 * 7
	}

	return Config{
		AccessTokenExpiry:  accessTokenExpiry,
		RefreshTokenExpiry: refreshTokenExpiry,
		JWTSecret:          []byte(secret),
	}
}

type AccessTokenFactory struct{}

func (f *AccessTokenFactory) CreateToken() Token {
	return &AccessToken{}
}
