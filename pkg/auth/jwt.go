package auth

import (
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateAccessToken(tokenString string, config utils.ServerConfig) (*CurrentUser, results.Result) {
	currentUser := &CurrentUser{}
	token, err := jwt.ParseWithClaims(tokenString, currentUser, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil || !token.Valid {
		return nil, results.NewUnauthorizedError("invalid access token")
	}

	return currentUser, results.NewResultOk()
}

func ValidateRefreshToken(tokenString string, config utils.ServerConfig) (*RefreshToken, results.Result) {
	currentUser := &RefreshToken{}
	token, err := jwt.ParseWithClaims(tokenString, currentUser, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET_REFRESH), nil
	})

	if err != nil || !token.Valid {
		return nil, results.NewUnauthorizedError("invalid access token")
	}

	return currentUser, results.NewResultOk()
}

func GenerateAccessToken(email string, ID uuid.UUID, config utils.ServerConfig) (string, results.Result) {
	expirationTime := time.Now().Add(60 * time.Minute) // Token expires in 5 minutes
	currentUser := &CurrentUser{
		Email: email,
		ID:    ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, currentUser)

	signedToken, err := token.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return "", results.NewBadRequestError(err.Error())
	}

	return signedToken, results.NewResultOk()
}

func GenerateRefreshToken(ID uuid.UUID, sessionID uuid.UUID, config utils.ServerConfig) (string, results.Result) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // Token expires in 7 days
	currentUser := &RefreshToken{	
		SessionID: sessionID,
		ID:    ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, currentUser)

	signedToken, err := token.SignedString([]byte(config.JWT_SECRET_REFRESH))

	if err != nil {
		return "", results.NewUnauthorizedError(err.Error())
	}
	
	return signedToken, results.NewResultOk()
}
