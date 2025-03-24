package utils

import (
	"errors"
	"os"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

// ValidateAccessToken validates the access token
var accessKey string = os.Getenv("JWT_ACCESS_SECRETE")
var refreshKey string = os.Getenv("JWT_REFRESH_SECRETE")

func ValidateAccessToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return accessKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	return claims, nil
}

// ValidateRefreshToken validates the refresh token
func ValidateRefreshToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	return claims, nil
}
