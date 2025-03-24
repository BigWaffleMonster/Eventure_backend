package service

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/models"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/views"
	"github.com/google/uuid"
)

type EventService struct {
	Repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{Repo: repo}
}

func (s *EventService) Create(data views.EventInfo) (string, error) {
	event := models.Event{
		ID: uuid.New(),
		OwnerID: uuid.New(),//TODO: Get executing author id from context
		Title: data.Title,
		Description: data.Description,
		Location: data.Location,
		Private: data.Private,
		StartDate: data.StartDate,
		EndDate: data.EndDate,//TODO: check for dates
		CategoryID: data.CategoryID,
	}
	
	err := s.Repo.Create(&event)
	if err != nil {
		return "", err
	}

	return "Successfully created!", nil
}