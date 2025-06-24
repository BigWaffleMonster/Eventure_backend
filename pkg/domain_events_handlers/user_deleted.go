package domain_events_handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
)

type userDeletedDomainEventHandler struct{
	EventRepo event.EventRepository
	ParticipantRepo participant.ParticipantRepository
}

func NewUserDeletedDomainEventHandler(eRepo event.EventRepository, pRepo participant.ParticipantRepository) domain_events_abstractions.DomainEventHandler{
	return &userDeletedDomainEventHandler{
		EventRepo: eRepo,
		ParticipantRepo: pRepo,
	}
}

func (h * userDeletedDomainEventHandler) Handle(domainEventData *domain_events_abstractions.DomainEventData) error {

	if domainEventData.Type != "UserDeletedDomainEvent" {
		return fmt.Errorf("failed to handler event. event type is incorrect")
	}

	var domainEvent domain_events.UserDeletedDomainEvent
	err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
	if err != nil {
		log.Print(err)
	}

	var events *[]event.Event

	events, err = h.EventRepo.GetOwnedCollection(domainEvent.UserID)
	if err != nil {
		return err
	}

	for _, e := range *events {
		err = h.EventRepo.Remove(e.ID)

		if err != nil {
			return err
		}

		var participants *[]participant.Participant

		participants, err = h.ParticipantRepo.GetCollection(e.ID)
		if err != nil {
			return err
		}

		for _, p := range *participants {
			err = h.ParticipantRepo.Remove(p.ID)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (h * userDeletedDomainEventHandler) IsTypeOf(domainEventData *domain_events_abstractions.DomainEventData) bool {
	return domainEventData.Type == "UserDeletedDomainEvent"
}