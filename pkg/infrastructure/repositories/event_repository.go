package repositories

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{DB: db}
}

func (r *EventRepository) Create(event *models.Event) error {
	return r.DB.Create(event).Error
}

func (r *EventRepository) Delete(event *models.Event) error {
	return r.DB.Delete(event).Error
}

func (r *EventRepository) Update(event *models.Event) error {
	return r.DB.Save(event).Error
}

func (r *EventRepository) GetCollection() error{
	var event models.Event
	return r.DB.Find(&event).Error
}
