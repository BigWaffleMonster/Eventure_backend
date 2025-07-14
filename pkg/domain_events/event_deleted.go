package domain_events

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type EventDeletedDomainEvent struct{
	EventID uuid.UUID
}

func NewEventDeletedDomainEvent(eventId uuid.UUID) (*domain_events_abstractions.DomainEventData, results.Result){
	event := EventDeletedDomainEvent{
		EventID: eventId,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, results.NewInternalError(err.Error())
    }

	domainEventData := NewDomainEventData("EventDeletedDomainEvent", string(b))

	return &domainEventData, results.NewResultOk()
}
