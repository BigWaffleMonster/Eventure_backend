package participant

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_definitions"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/utils/enums"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

//----------------------------------------------------------------------------------------------------------------------

type eventDeletedHandler struct{
	Repo ParticipantRepository
}

func NewEventDeletedHandler(repo ParticipantRepository) interfaces.DomainEventHandler{
	return &eventDeletedHandler{
		Repo: repo,
	}
}

func (h * eventDeletedHandler) Handle(domainEventData *domain_events_base.DomainEventData) results.Result {

	if domainEventData.Type != enums.EventDeleted {
		return results.NewInvalidDomainTypeError(domainEventData.Type.String(), enums.EventDeleted.String())
	}

	var domainEvent domain_events_definitions.EventDeleted
	err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
    if err != nil {
        return results.NewInternalError(err.Error())
    }

	var participants *[]Participant

	participants, result := h.Repo.GetCollectionByExpression("event_id = ?", domainEvent.EventID)
	if result.IsFailed {
		return result
	}

	for _, p := range *participants {
		result = h.Repo.Delete(p.ID)

		if result.IsFailed {
			return result
		}
	}

	return results.NewResultOk()
}

func (h * eventDeletedHandler) IsTypeOf(domainEventData *domain_events_base.DomainEventData) bool {
	return domainEventData.Type == enums.EventDeleted
}

//----------------------------------------------------------------------------------------------------------------------

type userCanVisitEventHandler struct{
	Repo ParticipantRepository
}

func NewUserCanVisitEventHandler(repo ParticipantRepository) interfaces.DomainEventHandler{
	return &userCanVisitEventHandler{
		Repo: repo,
	}
}

func (h * userCanVisitEventHandler) Handle(domainEventData *domain_events_base.DomainEventData) results.Result {

	if domainEventData.Type != enums.UserCanVisitEvent {
		return results.NewInvalidDomainTypeError(domainEventData.Type.String(), enums.UserCanVisitEvent.String())
	}

	var domainEvent domain_events_definitions.UserCanVisitEvent
	err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
    if err != nil {
        return results.NewInternalError(err.Error())
    }

	participant := Participant{
		ID: uuid.New(),
		UserID: domainEvent.UserID,
		EventID: domainEvent.EventID,
		Status: domainEvent.Status,
		Ticket: uuid.NewString(),
	}
	
	return h.Repo.Create(&participant)
}

func (h * userCanVisitEventHandler) IsTypeOf(domainEventData *domain_events_base.DomainEventData) bool {
	return domainEventData.Type == enums.UserCanVisitEvent
}