package auth

import (
	"context"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/config"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateAccessToken(ctx context.Context, tokenString string) (*CurrentUser, results.Result) {
	currentUser := &CurrentUser{}
	token, err := jwt.ParseWithClaims(tokenString, currentUser, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTSecret()), nil
	})

	if err != nil || !token.Valid {
		return nil, results.NewUnauthorizedError("invalid access token")
	}

	return currentUser, results.NewResultOk()
}

func ValidateRefreshToken(ctx context.Context, tokenString string) (*RefreshToken, results.Result) {
	currentUser := &RefreshToken{}
	token, err := jwt.ParseWithClaims(tokenString, currentUser, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTRefreshSecret()), nil
	})

	if err != nil || !token.Valid {
		return nil, results.NewUnauthorizedError("invalid access token")
	}

	return currentUser, results.NewResultOk()
}

func GenerateAccessToken(ctx context.Context, email string, ID uuid.UUID) (string, results.Result) {
	expirationTime := time.Now().Add(60 * time.Minute)
	currentUser := &CurrentUser{
		Email: email,
		ID:    ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, currentUser)

	signedToken, err := token.SignedString([]byte(config.GetJWTSecret()))

	if err != nil {
		return "", results.NewBadRequestError(err.Error())
	}

	return signedToken, results.NewResultOk()
}

func GenerateRefreshToken(ctx context.Context, ID uuid.UUID, sessionID uuid.UUID) (string, results.Result) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // Token expires in 7 days
	currentUser := &RefreshToken{	
		SessionID: sessionID,
		ID:    ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, currentUser)

	signedToken, err := token.SignedString([]byte(config.GetJWTRefreshSecret()))

	if err != nil {
		return "", results.NewUnauthorizedError(err.Error())
	}
	
	return signedToken, results.NewResultOk()
}
