package event

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/helpers"
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
	Repository EventRepository
	DomainEventBus domain_events_abstractions.DomainEventBus
}

func NewEventService(repository EventRepository, eventBus domain_events_abstractions.DomainEventBus) EventService {
	return &eventService{
		Repository: repository,
		DomainEventBus: eventBus,
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
	
	return s.Repository.Create(&event)
}

func (s *eventService) Update(id uuid.UUID, data *EventInput) results.Result {
	event, result := s.Repository.GetByID(id)

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

	return s.Repository.Update(event)
}

func (s *eventService) Delete(id uuid.UUID) results.Result {
	domainEventData, result := domain_events.NewEventDeletedDomainEvent(id)

	if result.IsFailed {
		return result
	}

	result = s.DomainEventBus.AddToStore(domainEventData)

	if result.IsFailed {
		return result
	}

	return s.Repository.Delete(id)
}

func (s *eventService) GetByID(id uuid.UUID) (*EventView, results.Result) {
	event, result := s.Repository.GetByID(id)
	if result.IsFailed {
		return nil, result
	}

	var eventView EventView

	copier.Copy(&eventView, &event)

	return &eventView, results.NewResultOk()
}

func (s *eventService) GetCollection() (*[]EventView, results.Result) {
	var events *[]Event

	events, result := s.Repository.GetCollection()
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

	events, result := s.Repository.GetOwnedCollection(currentUser.ID)
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