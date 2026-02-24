package schema

import "github.com/google/uuid"

type Category struct {
	ID    uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	Title string    `json:"title"`
}
