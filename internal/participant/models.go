package participant

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserID      uuid.UUID `gorm:"type:uuid;index"`
	EventID     uuid.UUID `gorm:"type:uuid;index"`
	Status      string    `sql:"type:ENUM('Yes', 'No', 'Maybe')"`
	Ticket      string
	DateCreated time.Time `gorm:"autoCreateTime"`
}

// @description participant input
type ParticipantInput struct {
	EventID     *uuid.UUID `json:"eventId" example:"09149ADB-CA29-401E-B9E9-06578A0A716C"`
	Status      *string    `json:"status" example:"Yes"`
}

// @description participant view
type ParticipantView struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userId"`
	EventID     uuid.UUID `json:"eventId"`
	Status      string    `json:"status"`
	Ticket      string    `json:"ticket"`
	DateCreated time.Time `json:"createdAt"`
}
