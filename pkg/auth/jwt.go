package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// ValidateAccessToken validates the access token
var accessKey string = os.Getenv("JWT_ACCESS_SECRET")
var refreshKey string = os.Getenv("JWT_REFRESH_SECRET")

func ValidateAccessToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return accessKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	return claims, nil
}

// ValidateRefreshToken validates the refresh token
func ValidateRefreshToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	return claims, nil
}

func GenerateAccessToken(email string, ID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute) // Token expires in 5 minutes
	claims := &Claims{
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
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // Token expires in 7 days
	claims := &Claims{
		Email: email,
		ID:    ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(os.Getenv(refreshKey))
}
