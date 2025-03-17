package models

import (
	"github.com/google/uuid"
)

type Notification struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uuid.UUID
	Title  string
}
