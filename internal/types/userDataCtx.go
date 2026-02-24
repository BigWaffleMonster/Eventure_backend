package types

import "github.com/google/uuid"

type UserDataCtx struct {
	UserID uuid.UUID
	Email  string
	Login  string
}
