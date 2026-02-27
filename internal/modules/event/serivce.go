package event

import (
	"net/http"
	"time"

	t "github.com/BigWaffleMonster/Eventure_backend/internal/types"
	global_utils "github.com/BigWaffleMonster/Eventure_backend/internal/utils"
	"github.com/google/uuid"
)

type EventService struct {
	repo *EventRepository
}

func NewEventService(repo *EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(req *CreateEventRequest, userDataCtx *t.UserDataCtx) (*EventResponse, error) {
	if req.EndDate != nil {
		if req.StartDate.After(*req.EndDate) || req.StartDate.Equal(*req.EndDate) {
			return nil, global_utils.NewAppError(http.StatusBadRequest, "дата начала должна быть раньше даты окончания")
		}
	}

	if req.StartDate.Before(time.Now()) {
		return nil, global_utils.NewAppError(http.StatusBadRequest, "дата начала должна быть в будущем")
	}

	category, err := s.repo.GetCategoryForEventByID(req.CategoryID)
	if err != nil {
		return nil, err
	}

	event, err := s.repo.CreateEvent(req, userDataCtx.UserID, category.ID)
	if err != nil {
		return nil, err
	}

	careatedEvent := EventResponse{
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

func (s *EventService) GetEventByID(eventID uuid.UUID) (*EventResponse, error) {
	event, err := s.repo.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	eventDTO := EventResponse{
		ID:          eventID,
		Title:       event.Title,
		Description: event.Description,
		Capacity:    *event.Capacity,
		MaxCapacity: event.MaxCapacity,
		Location:    event.Location,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		DateCreated: event.DateCreated,
		DateUpdated: event.DateUpdated,

		Category: CategoryResponse(event.Category),
		Owner:    OwnerResponse{ID: event.Owner.ID, Login: event.Owner.Login, Email: event.Owner.Email},
	}

	return &eventDTO, nil
}
