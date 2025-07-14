package domain_events_handlers

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
)

const eventDeletedDomainType = "EventDeletedDomainEvent"

type eventDeletedDomainEventHandler struct{
	ParticipantRepo participant.ParticipantRepository
}

func NewEventDeletedDomainEventHandler(pRepo participant.ParticipantRepository) domain_events_abstractions.DomainEventHandler{
	return &eventDeletedDomainEventHandler{
		ParticipantRepo: pRepo,
	}
}

func (h * eventDeletedDomainEventHandler) Handle(domainEventData *domain_events_abstractions.DomainEventData) results.Result {

	if domainEventData.Type != eventDeletedDomainType {
		return results.NewInvalidDomainTypeError(domainEventData.Type, eventDeletedDomainType)
	}

	var domainEvent domain_events.EventDeletedDomainEvent
	err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
    if err != nil {
        return results.NewInternalError(err.Error())
    }

	var participants *[]participant.Participant

	participants, result := h.ParticipantRepo.GetCollection(domainEvent.EventID)
	if result.IsFailed {
		return result
	}

	for _, p := range *participants {
		result = h.ParticipantRepo.Delete(p.ID)

		if result.IsFailed {
			return result
		}
	}

	return results.NewResultOk()
}

func (h * eventDeletedDomainEventHandler) IsTypeOf(domainEventData *domain_events_abstractions.DomainEventData) bool {
	return domainEventData.Type == eventDeletedDomainType
}