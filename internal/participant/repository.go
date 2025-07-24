package participant

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"gorm.io/gorm"
)

func NewParticipantRepository(db *gorm.DB) repository.Repository[Participant] {
	return repository.NewRepository[Participant](db)
}