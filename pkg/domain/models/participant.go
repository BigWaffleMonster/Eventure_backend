package models

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	UserID      uuid.UUID `gorm:"primaryKey;type:uuid"`
	EventID     uuid.UUID `gorm:"primaryKey;type:uuid"`
	Status      string    `sql:"type:ENUM('Yes', 'No', 'Maybe')"`
	HasAccess   bool
	Ticket      string
	DateCreated time.Time `gorm:"autoCreateTime"`
}
