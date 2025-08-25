package category

import "github.com/google/uuid"

type Category struct {
	ID    uuid.UUID `gorm:"primaryKey;type:uuid"`
	Title string
}

type CategoryView struct {
	ID    uuid.UUID   `json:"id"`
	Title string      `json:"title"`
}
