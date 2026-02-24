package event

import (
	"errors"
	"time"

	t "github.com/BigWaffleMonster/Eventure_backend/internal/types"
)

type EventService struct {
	repo *EventRepository
}

func NewEventService(repo *EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(req *CreateEventRequest, userDataCtx *t.UserDataCtx) (*CreateEventResponse, error) {
	if req.EndDate != nil {
		if req.StartDate.After(*req.EndDate) || req.StartDate.Equal(*req.EndDate) {
			return nil, errors.New("дата начала должна быть раньше даты окончания")
		}
	}

	if req.StartDate.Before(time.Now()) {
		return nil, errors.New("дата начала должна быть в будущем")
	}

	category, err := s.repo.GetCategoryForEventByID(req.CategoryID)
	if err != nil {
		return nil, err
	}

	event, err := s.repo.CreateEvent(req, userDataCtx.UserID, category.ID)
	if err != nil {
		return nil, err
	}

	careatedEvent := CreateEventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Capacity:    0,
		MaxCapacity: event.MaxCapacity,
		Location:    event.Location,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,

		DateCreated: time.Now(),
		DateUpdated: time.Now(),

		Category: CategoryResponse{ID: category.ID, Title: category.Title},
		Owner:    OwnerResponse{ID: userDataCtx.UserID, Login: userDataCtx.Login, Email: userDataCtx.Email},
	}

	return &careatedEvent, nil
}
