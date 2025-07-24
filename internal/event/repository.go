package event

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"gorm.io/gorm"
)

func NewEventRepository(db *gorm.DB) repository.Repository[Event] {
	return repository.NewRepository[Event](db)
}