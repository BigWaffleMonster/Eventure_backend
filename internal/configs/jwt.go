package configs

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTConfig struct {
	AccessSecretKey  string
	RefreshSecretKey string
	AccessTokenExp   time.Duration
	RefreshTokenExp  time.Duration
	Issuer           string
}

func InitJWTConfig() *JWTConfig {
	return &JWTConfig{
		AccessSecretKey:  getEnv("JWT_ACCESS_SECRET", "your-super-secret-key-change-in-prod"),
		RefreshSecretKey: getEnv("JWT_REFRESH_SECRET", "your-super-secret-key-change-in-prod"),
		AccessTokenExp:   15 * time.Minute,
		RefreshTokenExp:  7 * 24 * time.Hour,
		Issuer:           "eventure-api",
	}
}

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Login  string    `json:"login"`
	jwt.RegisteredClaims
}
