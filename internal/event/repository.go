package event

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository interface{
    Create(event *Event)  error
	Update(event *Event) error
	Delete(ID uuid.UUID) error
	GetById(ID uuid.UUID) (*Event, error)
	GetCollection() ([]Event, error)
}

type EventRepositoryImp struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImp {
	return &EventRepositoryImp{DB: db}
}

func (repository EventRepositoryImp) Create(event *Event) error {
	return repository.DB.Create(event).Error
}

func (repository *EventRepositoryImp) Delete(ID uuid.UUID) error {
	var event Event
	return repository.DB.Where("id = ?", ID).Delete(&event).Error
}

func (repository *EventRepositoryImp) Update(event *Event) error {
	return repository.DB.Save(event).Error
}

func (repository *EventRepositoryImp) GetById(ID uuid.UUID) (*Event, error) {
	var event Event
	result := repository.DB.Where("id = ?", ID).First(&event)
	if result.Error != nil {
		return nil, result.Error
	}
	return &event, nil
}

func (repository*EventRepositoryImp) GetCollection() ([]Event, error){
	var events []Event

	result := repository.DB.Find(&events)

	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}