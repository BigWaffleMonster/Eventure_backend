package domain_events_handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
)

type eventDeletedDomainEventHandler struct{
	ParticipantRepo participant.ParticipantRepository
}

func NewEventDeletedDomainEventHandler(pRepo participant.ParticipantRepository) domain_events_abstractions.DomainEventHandler{
	return &eventDeletedDomainEventHandler{
		ParticipantRepo: pRepo,
	}
}

func (h * eventDeletedDomainEventHandler) Handle(domainEventData *domain_events_abstractions.DomainEventData) error {

	if domainEventData.Type != "EventDeletedDomainEvent" {
		return fmt.Errorf("Failed")
	}

	var domainEvent domain_events.EventDeletedDomainEvent
	err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
	if err != nil {
		log.Print(err)
	}

	var participants *[]participant.Participant

	participants, err = h.ParticipantRepo.GetCollection(domainEvent.EventID)
	if err != nil {
		return err
	}

	for _, p := range *participants {
		err = h.ParticipantRepo.Remove(p.ID)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h * eventDeletedDomainEventHandler) IsTypeOf(domainEventData *domain_events_abstractions.DomainEventData) bool {
	return domainEventData.Type == "EventDeletedDomainEvent"
}