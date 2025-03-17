package models

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	UserID      uuid.UUID `gorm:"primaryKey"`
	EventID     uuid.UUID `gorm:"primaryKey"`
	Status      string    `sql:"type:ENUM('Yes', 'No', 'Maybe')"`
	HasAccess   bool
	Ticket      string
	DateCreated time.Time `gorm:"autoCreateTime"`
}
