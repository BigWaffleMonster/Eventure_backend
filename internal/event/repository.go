package event

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"gorm.io/gorm"
)

type EventRepository interface {
	interfaces.IBaseRepository[Event]
}

type eventRepository struct {
	repository.BaseRepository[Event]
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{
		repository.BaseRepository[Event]{DB: db},
	}
}