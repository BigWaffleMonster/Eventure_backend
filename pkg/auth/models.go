package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	Email string
	ID    uuid.UUID
	jwt.RegisteredClaims
}
