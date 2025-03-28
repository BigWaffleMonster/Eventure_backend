package event

import (
	"errors"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/helpers"
	"github.com/google/uuid"
	"github.com/stroiman/go-automapper"
)

type IEventService interface{
    Create(data EventInput) error
	Update(ID uuid.UUID, data EventInput) error
	Delete(ID uuid.UUID) error
	GetById(ID uuid.UUID) (*EventView, error)
	GetCollection() ([]EventView, error)
}

type EventService struct {
	Repository EventRepository
}

func NewEventService(repository EventRepository) *EventService {
	return &EventService{Repository: repository}
}

func (service *EventService) Create(data EventInput) (error) {

	if data.EndDate.Before(data.StartDate) {
		return errors.New("date errors")
	}

	event := Event{
		ID: uuid.New(),
		OwnerID: uuid.New(),//TODO: add un middleware and get user from it
		Title: data.Title,
		Description: data.Description,
		Location: data.Location,
		Private: data.Private,
		StartDate: data.StartDate,
		EndDate: data.EndDate,
		CategoryID: data.CategoryID,
	}
	
	return service.Repository.Create(&event)
}

func (service *EventService) Update(ID uuid.UUID, data EventInput) (error) {
	event, err := service.Repository.GetById(ID)
	if err != nil {
		return err
	}

	event.Title = data.Title
	event.Description = data.Description
	event.Location = data.Location
	event.Private = data.Private
	event.StartDate = data.StartDate
	event.EndDate = data.EndDate
	event.CategoryID = data.CategoryID
	event.MaxQtyParticipants = data.MaxQtyParticipants

	return service.Repository.Update(event)
}

func (service *EventService) Delete(ID uuid.UUID) (error) {
	return service.Repository.Delete(ID)
}

func (service *EventService) GetById(ID uuid.UUID) (*EventView, error) {
	event, err := service.Repository.GetById(ID)
	if err != nil {
		return nil, err
	}

	var eventView EventView

	automapper.Map(event, eventView)

	return &eventView, nil
}

func (service *EventService) GetCollection() ([]EventView, error) {
	events, err := service.Repository.GetCollection()
	if err != nil {
		return []EventView{}, err
	}

	views := helpers.MapArray(events, func(event Event) EventView {
		var eventView EventView
		automapper.Map(event, eventView)
		return eventView
	})


	return views, nil
}