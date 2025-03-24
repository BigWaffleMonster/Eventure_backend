package services

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/application/views"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain/models"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/infrastructure/repositories"
	"github.com/google/uuid"
)

type EventService struct {
	Repo *repositories.EventRepository
}

func NewEventService(repo *repositories.EventRepository) *EventService {
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