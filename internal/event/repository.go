package event

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository interface{
    Create(event *Event) results.Result
	Update(event *Event) results.Result
	Delete(id uuid.UUID) results.Result
	GetByID(id uuid.UUID) (*Event, results.Result)
	GetCollection() (*[]Event, results.Result)
	GetOwnedCollection(currentUserID uuid.UUID) (*[]Event, results.Result)
}

type eventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{DB: db}
}

func (r *eventRepository) Create(event *Event) results.Result {
	err := r.DB.Create(event).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *eventRepository) Delete(id uuid.UUID) results.Result {
	var event Event

	err := r.DB.Where("id = ?", id).Delete(&event).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *eventRepository) Update(event *Event) results.Result {
	err := r.DB.Save(event).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *eventRepository) GetByID(id uuid.UUID) (*Event, results.Result) {
	var event Event

	result := r.DB.First(&event, "id = ?", id)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &event, results.NewResultOk()
}

func (r *eventRepository) GetCollection() (*[]Event, results.Result){
	var events []Event

	result := r.DB.Find(&events)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &events, results.NewResultOk()
}

func (r *eventRepository) GetOwnedCollection(currentUserID uuid.UUID) (*[]Event, results.Result){
	var events []Event

	result := r.DB.Find(&events, "owner_id = ?", currentUserID)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &events, results.NewResultOk()
}