package domain_events

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/google/uuid"
)

type EventDeletedDomainEvent struct{
	EventID uuid.UUID
}

func NewEventDeletedDomainEvent(eventId uuid.UUID) (*domain_events_abstractions.DomainEventData, error){
	event := EventDeletedDomainEvent{
		EventID: eventId,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, err
    }

	domainEventData := NewDomainEventData("EventDeletedDomainEvent", string(b))

	return &domainEventData, nil
}

type UserDeletedDomainEvent struct{
	UserID uuid.UUID
}

func NewUserDeletedDomainEvent(userID uuid.UUID) (*domain_events_abstractions.DomainEventData, error){
	event := UserDeletedDomainEvent{
		UserID: userID,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, err
    }

	domainEventData := NewDomainEventData("UserDeletedDomainEvent", string(b))

	return &domainEventData, nil
}

type ParticipantCreatedDomainEvent struct{
	EventID uuid.UUID
	UserID uuid.UUID
	Status string
	Ticket string 
}

func NewParticipantCreatedDomainEvent(eventId uuid.UUID, userID uuid.UUID, status string, ticket string) (*domain_events_abstractions.DomainEventData, error){
	event := ParticipantCreatedDomainEvent{
		EventID: eventId,
		UserID: userID,
		Status: status,
		Ticket: ticket,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, err
    }

	domainEventData := NewDomainEventData("ParticipantCreatedDomainEvent", string(b))

	return &domainEventData, nil
}

func NewDomainEventData(eventType string, content string) domain_events_abstractions.DomainEventData{
	return domain_events_abstractions.DomainEventData{
		ID: uuid.New(),
		Type: eventType,
		Content: content,
	}
}