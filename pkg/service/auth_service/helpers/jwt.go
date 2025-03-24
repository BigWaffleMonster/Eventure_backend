package helpers

import (
	"os"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var accessKey string = os.Getenv("JWT_ACCESS_SECRETE")
var refreshKey string = os.Getenv("JWT_REFRESH_SECRETE")

func GenerateAccessToken(email string, ID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute) // Token expires in 5 minutes
	claims := &models.Claims{
		Email: email,
		ID:    ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessKey)
}

func GenerateRefreshToken(email string, ID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // Token expires in 5 minutes
	claims := &models.Claims{
		Email: email,
		ID:    ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(os.Getenv(refreshKey))
}
