package event

import (
	"errors"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/helpers"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type EventService interface{
    Create(data *EventInput) error
	Update(id uuid.UUID, data *EventInput) error
	Remove(id uuid.UUID) error
	GetByID(id uuid.UUID) (*EventView, error)
	GetCollection() (*[]EventView, error)
}

type eventService struct {
	Repository EventRepository
}

func NewEventService(repository EventRepository) EventService {
	return &eventService{Repository: repository}
}

func (s *eventService) Create(data *EventInput) (error) {

	if data.EndDate.Before(*data.StartDate) {
		return errors.New("date errors")
	}

	event := Event{
		ID: uuid.New(),
		OwnerID: uuid.New(),//TODO: add un middleware and get user from it
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

func (s *eventService) Update(id uuid.UUID, data *EventInput) (error) {
	event, err := s.Repository.GetByID(id)
	if err != nil {
		return err
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

func (s *eventService) Remove(id uuid.UUID) (error) {
	return s.Repository.Remove(id)
}

func (s *eventService) GetByID(id uuid.UUID) (*EventView, error) {
	event, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	var eventView EventView

	copier.Copy(&eventView, &event)

	return &eventView, nil
}

func (s *eventService) GetCollection() (*[]EventView, error) {
	var events *[]Event

	events, err := s.Repository.GetCollection()
	if err != nil {
		return nil, err
	}

	views := helpers.MapArray(events, func(event Event) EventView {
		var eventView EventView
		copier.Copy(&eventView, &event)
		return eventView
	})


	return views, nil
}