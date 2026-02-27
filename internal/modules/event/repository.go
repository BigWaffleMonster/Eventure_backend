package event

import (
	"errors"
	"net/http"
	"time"

	schema "github.com/BigWaffleMonster/Eventure_backend/internal/db/schema"
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

func (r *EventRepository) GetEvents() {}

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

func (r *EventRepository) GetUserCreatedEvents(userID string) {}

func (r *EventRepository) GetUserParticipatingEvents(userID string) {}

func (r *EventRepository) CreateEvent(event *CreateEventRequest, userID uuid.UUID, categoryID uuid.UUID) (*schema.Event, error) {
	newEvent := &schema.Event{
		ID:          uuid.New(),
		Title:       event.Title,
		Description: event.Description,
		MaxCapacity: event.MaxCapacity,
		Location:    event.Location,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,

		DateCreated: time.Now(),
		DateUpdated: time.Now(),

		CategoryID: &event.CategoryID,
		OwnerID:    userID,
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
