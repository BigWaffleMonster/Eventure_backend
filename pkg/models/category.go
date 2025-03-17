package models

import (
	"github.com/google/uuid"
)

type Category struct {
	ID    uuid.UUID `gorm:"primaryKey"`
	Title string
}
