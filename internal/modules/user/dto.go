package user

import "github.com/google/uuid"

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Login string    `json:"login"`
	Email string    `json:"email"`
}
