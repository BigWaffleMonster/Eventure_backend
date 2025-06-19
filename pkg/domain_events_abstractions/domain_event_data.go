package domain_events_abstractions

import (
	"time"

	"github.com/google/uuid"
)

const (
	NewEventCreatedDomainEvent = "NewEventCreatedDomainEvent"
	UserDeletedDomainEvent = "UserDeletedDomainEvent"
	EventDeletedDomainEvent = "EventDeletedDomainEvent"
	NewParticipantDomainEvent = "NewParticipantDomainEvent"
	UserRegisteredDomainEvent = "UserRegisteredDomainEvent"
)

type DomainEventData struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid"`
	Type string 
	Content string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}