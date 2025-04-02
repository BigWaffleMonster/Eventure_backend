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

type RegisterInput struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type LoginInput struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type RefreshInput struct {
	RefreshToken string `json:"refresh_token"`
}
