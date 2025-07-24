package category

import "github.com/google/uuid"

type Category struct {
	ID    uuid.UUID `gorm:"unique;primaryKey;autoIncrement"`
	Title string
}

type CategoryView struct {
	ID    uuid.UUID   `json:"id"`
	Title string `json:"title"`
}
