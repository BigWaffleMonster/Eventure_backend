package event

import (
	"fmt"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_definitions"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type EventService interface{
    Create(data *EventInput, currentUser *auth.CurrentUser) results.Result
	Update(id uuid.UUID, data *EventInput) results.Result
	Delete(id uuid.UUID) results.Result
	GetByID(id uuid.UUID) (*EventView, results.Result)
	GetCollection() (*[]EventView, results.Result)
	GetOwnedCollection(currentUser *auth.CurrentUser) (*[]EventView, results.Result)
}

type eventService struct {
	UOF UnitOfWork
}

func NewEventService(UOF UnitOfWork) EventService {
	return &eventService{
		UOF: UOF,
	}
}

func (s *eventService) Create(data *EventInput, currentUser *auth.CurrentUser) results.Result {

	if data.EndDate.Before(*data.StartDate) {
		return results.NewBadRequestError("End date is defore start date")
	}

	event := Event{
		ID: uuid.New(),
		OwnerID: currentUser.ID,
		MaxQtyParticipants: *data.MaxQtyParticipants,
		Title: *data.Title,
		Description: *data.Description,
		Location: *data.Location,
		Private: *data.Private,
		StartDate: *data.StartDate,
		EndDate: *data.EndDate,
		CategoryID: *data.CategoryID,
	}
	
	return s.UOF.Repository().Create(&event)
}

func (s *eventService) Update(id uuid.UUID, data *EventInput) results.Result {
	repository := s.UOF.Repository()

	event, result := repository.GetByID(id)

	if result.IsFailed {
		return result
	}

	if data.Title != nil {
		event.Title = *data.Title
	}

	if data.Description != nil {
		event.Description = *data.Description
	}

	if data.Location != nil {
		event.Location = *data.Location
	}

	if data.Private != nil {
		event.Private = *data.Private
	}

	if data.StartDate != nil {
		event.StartDate = *data.StartDate
	}

	if data.EndDate != nil {
		event.EndDate = *data.EndDate
	}

	if data.CategoryID != nil {
		event.CategoryID = *data.CategoryID
	}

	if data.MaxQtyParticipants != nil {
		event.MaxQtyParticipants = *data.MaxQtyParticipants
	}

	return repository.Update(event)
}

func (s *eventService) Delete(id uuid.UUID) results.Result {
	return s.UOF.RunInTx(
		NewEventRepository,
		func(repo EventRepository, store interfaces.DomainEventStore) results.Result{
		domainEventData, result := domain_events_definitions.NewEventDeleted(id)

		if result.IsFailed {
			return result
		}

		result = store.AddToStore(domainEventData)

		if result.IsFailed {
			return result
		}
		
		return repo.Delete(id)
	})
}

func (s *eventService) GetByID(id uuid.UUID) (*EventView, results.Result) {
	event, result := s.UOF.Repository().GetByID(id)
	if result.IsFailed {
		return nil, result
	}

	var eventView EventView

	copier.Copy(&eventView, &event)

	return &eventView, results.NewResultOk()
}

func (s *eventService) GetCollection() (*[]EventView, results.Result) {
	var events *[]Event

	events, result := s.UOF.Repository().GetCollection()
	if result.IsFailed {
		return nil, result
	}

	views := helpers.MapArray(events, func(event Event) EventView {
		var eventView EventView
		copier.Copy(&eventView, &event)
		return eventView
	})

	return views, results.NewResultOk()
}

func (s *eventService) GetOwnedCollection(currentUser *auth.CurrentUser) (*[]EventView, results.Result) {
	var events *[]Event

	fmt.Println(currentUser.ID)

	events, result := s.UOF.Repository().GetCollectionByExpression("owner_id = ?", currentUser.ID)
	if result.IsFailed {
		return nil, result
	}

	views := helpers.MapArray(events, func(event Event) EventView {
		var eventView EventView
		copier.Copy(&eventView, &event)
		return eventView
	})

	return views, results.NewResultOk()
}