package participant

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Login string    `json:"login"`
	Email string    `json:"email"`
}

type ParticipantResponse struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	EventID     uuid.UUID `gorm:"type:uuid;index" json:"event_id"`
	DateCreated time.Time `gorm:"autoCreateTime" json:"date_created"`

	User UserResponse
}
