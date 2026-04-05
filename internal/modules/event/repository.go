package event

import (
	"errors"
	"net/http"
	"time"

	schema "github.com/BigWaffleMonster/Eventure_backend/internal/db/schema"
	"github.com/BigWaffleMonster/Eventure_backend/internal/types"
	"github.com/google/uuid"
	"gorm.io/gorm"

	global_utils "github.com/BigWaffleMonster/Eventure_backend/internal/utils"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) GetEvents(params *PaginationParams) ([]schema.Event, *int64, error) {
	var events_raw []schema.Event
	var total int64

	r.db.Model(&schema.Event{}).Count(&total)

	err := r.db.Preload("Category").Preload("Owner").Offset(params.Offset).Limit(params.Limit).Find(&events_raw).Error
	if err != nil {
		return nil, nil, global_utils.NewAppErrorWithErr(http.StatusInternalServerError, "Ошибка получения событий", err)
	}

	return events_raw, &total, nil
}

func (r *EventRepository) GetEventByID(id uuid.UUID) (*schema.Event, error) {
	var event schema.Event

	err := r.db.Preload("Category").Preload("Owner").First(&event, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, global_utils.ErrNotFound
		}
		return nil, global_utils.NewAppErrorWithErr(http.StatusInternalServerError, "Ошибка получения события", err)
	}

	return &event, nil
}

func (r *EventRepository) GetUserCreatedEvents(userID uuid.UUID) ([]schema.Event, error) {
	var events_raw []schema.Event

	err := r.db.Preload("Category").Preload("Owner").Where("owner_id = ?", userID).Find(&events_raw).Error
	if err != nil {
		return nil, global_utils.NewAppErrorWithErr(http.StatusInternalServerError, "Ошибка получения событий", err)
	}

	return events_raw, nil
}

func (r *EventRepository) GetUserParticipatingEvents(userID uuid.UUID) ([]schema.Participant, error) {
	var participants_raw []schema.Participant

	err := r.db.Preload("Event").Preload("User").Preload("Event.Category").Preload("Event.Owner").Where("user_id = ?", userID).Find(&participants_raw).Error
	if err != nil {
		return nil, global_utils.NewAppErrorWithErr(http.StatusInternalServerError, "Ошибка получения событий", err)
	}

	return participants_raw, nil
}

func (r *EventRepository) CreateEvent(event *CreateEventRequest, userID uuid.UUID, categoryID uuid.UUID, coverURL *string) (*schema.Event, error) {
	newEvent := &schema.Event{
		ID:          uuid.New(),
		Title:       event.Title,
		Description: event.Description,
		MaxCapacity: event.MaxCapacity,
		Location:    (schema.Location)(event.Location),
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,

		DateCreated: time.Now(),
		DateUpdated: time.Now(),

		CategoryID: event.CategoryID,
		OwnerID:    userID,

		Cover: coverURL,
	}

	if err := r.db.Create(newEvent).Error; err != nil {
		return nil, global_utils.NewAppErrorWithErr(http.StatusBadRequest, "Ошибка создания события", err)
	}

	return newEvent, nil
}

func (r *EventRepository) GetCategoryForEventByID(id uuid.UUID) (*schema.Category, error) {
	var category schema.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, global_utils.NewAppError(http.StatusBadRequest, "Категория не найдена")
		}
		return nil, global_utils.NewAppErrorWithErr(http.StatusBadRequest, "Ошибка поиска категории", err)
	}

	return &category, nil
}

func (r *EventRepository) UpdateEvent(eventID uuid.UUID, userData *types.UserDataCtx, data *UpdateEventRequest) error {
	var event schema.Event

	err := r.db.Where("id = ? AND owner_id = ?", eventID, userData.UserID).First(&event).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return global_utils.NewAppError(http.StatusBadRequest, "Событие не найдена")
		}
		return global_utils.NewAppErrorWithErr(http.StatusBadRequest, "Ошибка поиска события", err)
	}

	if data.Title != nil {
		event.Title = *data.Title
	}
	if data.Description != nil {
		event.Description = *data.Description
	}
	if data.Capacity != nil {
		event.Capacity = *data.Capacity
	}
	if data.MaxCapacity != nil {
		event.MaxCapacity = *data.MaxCapacity
	}
	if data.Location != nil {
		event.Location = (schema.Location)(event.Location)
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

	if err := r.db.Save(&event).Error; err != nil {
		return global_utils.NewAppErrorWithErr(http.StatusBadRequest, "Ошибка сохранения обновлений", err)
	}

	return nil
}

func (r *EventRepository) RemoveEvent(id uuid.UUID) error {
	err := r.db.Delete(&schema.Event{}, id).Error
	if err != nil {
		return global_utils.NewAppErrorWithErr(http.StatusBadRequest, "Ошибка поиска категории", err)
	}

	return nil
}
