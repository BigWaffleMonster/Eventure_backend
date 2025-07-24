package participant

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"gorm.io/gorm"
)

type ParticipantRepository interface {
	interfaces.IBaseRepository[Participant]
}

type participantRepository struct {
	repository.BaseRepository[Participant]
}

func NewParticipantRepository(db *gorm.DB) ParticipantRepository {
	return &participantRepository{
		repository.BaseRepository[Participant]{DB: db},
	}
}