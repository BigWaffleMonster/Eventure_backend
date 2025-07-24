package event

import (
	"encoding/json"
	"fmt"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_definitions"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/utils/enums"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
)

type userWantsToVisitEventHandler struct{
	UOF UnitOfWork
}

func NewUserWantsToVisitEventHandler(uof UnitOfWork) interfaces.DomainEventHandler{
	return &userWantsToVisitEventHandler{
		UOF: uof,
	}
}

func (h * userWantsToVisitEventHandler) Handle(domainEventData *domain_events_base.DomainEventData) results.Result {
	return h.UOF.RunInTx(
		NewEventRepository,
		func(repo EventRepository, store interfaces.DomainEventStore) results.Result{
			if domainEventData.Type != enums.UserWantsToVisitEvent {
				return results.NewInvalidDomainTypeError(domainEventData.Type.String(), enums.UserWantsToVisitEvent.String())
			}

			var domainEvent domain_events_definitions.UserWantsToVisitEvent
			err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
			if err != nil {
				return results.NewInternalError(err.Error())
			}

			var event *Event

			event, result := repo.GetByID(domainEvent.EventID)
			if result.IsFailed {
				return result
			}

			if event.MaxQtyParticipants > domainEvent.ActualQTYOfGuests{
				//TODO: Notify user
				fmt.Println("failed to create new participant, max QTY of participants reached")
				return results.NewResultOk()
			}

			result = repo.Update(event)

			if result.IsFailed {
				return result
			}

			userCanVisitEvent, result := domain_events_definitions.NewUserCanVisitEvent(
				domainEvent.EventID,
				domainEvent.UserID,
				domainEvent.Status,
			)

			if result.IsFailed {
				return result
			}

			return store.AddToStore(userCanVisitEvent)
		})
}

func (h * userWantsToVisitEventHandler) IsTypeOf(domainEventData *domain_events_base.DomainEventData) bool {
	return domainEventData.Type == enums.UserWantsToVisitEvent
}


//----------------------------------------------------------------------------------------------------------------------

type userDeletedHandler struct{
	UOF UnitOfWork
}

func NewUserDeletedHandler(uof UnitOfWork) interfaces.DomainEventHandler{
	return &userDeletedHandler{
		UOF: uof,
	}
}

func (h * userDeletedHandler) Handle(domainEventData *domain_events_base.DomainEventData) results.Result {
	return h.UOF.RunInTx(
		NewEventRepository,
		func(repo EventRepository, store interfaces.DomainEventStore) results.Result{
			if domainEventData.Type != enums.UserDeleted {
				return results.NewInvalidDomainTypeError(domainEventData.Type.String(), enums.UserDeleted.String())
			}

			var domainEvent domain_events_definitions.UserDeleted
			err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
			if err != nil {
				return results.NewInternalError(err.Error())
			}

			var events *[]Event

			events, result := repo.GetCollectionByExpression("owner_id = ?", domainEvent.UserID)
			if result.IsFailed {
				return result
			}

			for _, e := range *events {
				result = repo.Delete(e.ID)

				if result.IsFailed {
					return result
				}
				
				domainEventData, result := domain_events_definitions.NewEventDeleted(e.ID)

				if result.IsFailed {
					return result
				}

				result = store.AddToStore(domainEventData)

				if result.IsFailed {
					return result
				}
			}

			return results.NewResultOk()
	})
}

func (h * userDeletedHandler) IsTypeOf(domainEventData *domain_events_base.DomainEventData) bool {
	return domainEventData.Type == enums.UserDeleted
}