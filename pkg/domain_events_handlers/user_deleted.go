package domain_events_handlers

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
)

const userDeletedDomainType = "UserDeletedDomainEvent"

type userDeletedDomainEventHandler struct{
	EventRepo repository.Repository[event.Event]
	ParticipantRepo repository.Repository[participant.Participant]
}

func NewUserDeletedDomainEventHandler(
	eRepo repository.Repository[event.Event], 
	pRepo repository.Repository[participant.Participant]) domain_events_abstractions.DomainEventHandler{
	return &userDeletedDomainEventHandler{
		EventRepo: eRepo,
		ParticipantRepo: pRepo,
	}
}

func (h * userDeletedDomainEventHandler) Handle(domainEventData *domain_events_abstractions.DomainEventData) results.Result {

	if domainEventData.Type != userDeletedDomainType {
		return results.NewInvalidDomainTypeError(domainEventData.Type, userDeletedDomainType)
	}

	var domainEvent domain_events.UserDeletedDomainEvent
	err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
    if err != nil {
        return results.NewInternalError(err.Error())
    }

	var events *[]event.Event

	events, result := h.EventRepo.GetCollectionByExpression("owner_id = ?", domainEvent.UserID)
	if result.IsFailed {
		return result
	}

	for _, e := range *events {
		result = h.EventRepo.Delete(e.ID)

		if result.IsFailed {
			return result
		}

		var participants *[]participant.Participant

		participants, result = h.ParticipantRepo.GetCollectionByExpression("event_id = ?", e.ID)
		if result.IsFailed {
			return result
		}

		for _, p := range *participants {
			result = h.ParticipantRepo.Delete(p.ID)

			if result.IsFailed {
				return result
			}
		}
	}

	return results.NewResultOk()
}

func (h * userDeletedDomainEventHandler) IsTypeOf(domainEventData *domain_events_abstractions.DomainEventData) bool {
	return domainEventData.Type == userDeletedDomainType
}