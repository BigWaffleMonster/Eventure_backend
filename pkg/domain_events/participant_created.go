package domain_events

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type ParticipantCreatedDomainEvent struct{
	EventID uuid.UUID
	UserID uuid.UUID
	Status string
	Ticket string 
}

func NewParticipantCreatedDomainEvent(
	eventId uuid.UUID, 
	userID uuid.UUID, 
	status string, 
	ticket string) (*domain_events_abstractions.DomainEventData, results.Result){
		
	event := ParticipantCreatedDomainEvent{
		EventID: eventId,
		UserID: userID,
		Status: status,
		Ticket: ticket,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, results.NewInternalError(err.Error())
    }

	domainEventData := NewDomainEventData("ParticipantCreatedDomainEvent", string(b))

	return &domainEventData, results.NewResultOk()
}