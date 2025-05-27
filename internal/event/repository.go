package event

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository interface{
    Create(event *Event) error
	Update(event *Event) error
	Remove(id uuid.UUID) error
	GetByID(id uuid.UUID) (*Event, error)
	GetCollection() (*[]Event, error)
	GetOwnedCollection(currentUserID uuid.UUID) (*[]Event, error)
}

type eventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{DB: db}
}

func (r *eventRepository) Create(event *Event) error {
	return r.DB.Create(event).Error
}

func (r *eventRepository) Remove(id uuid.UUID) error {
	var event Event
	return r.DB.Where("id = ?", id).Delete(&event).Error
}

func (r *eventRepository) Update(event *Event) error {
	return r.DB.Save(event).Error
}

func (r *eventRepository) GetByID(id uuid.UUID) (*Event, error) {
	var event Event
	result := r.DB.First(&event, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &event, nil
}

func (r *eventRepository) GetCollection() (*[]Event, error){
	var events []Event

	result := r.DB.Find(&events)

	if result.Error != nil {
		return nil, result.Error
	}
	return &events, nil
}

func (r *eventRepository) GetOwnedCollection(currentUserID uuid.UUID) (*[]Event, error){
	var events []Event

	result := r.DB.Find(&events, "owner_id = ?", currentUserID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &events, nil
}