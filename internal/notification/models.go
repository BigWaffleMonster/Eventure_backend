package notification

import (
	"github.com/google/uuid"
)

type Notification struct {
	ID     uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserID uuid.UUID
	Title  string
}
