package participant

import (
	"github.com/google/uuid"
)

// "net/http"
// "time"

// "github.com/BigWaffleMonster/Eventure_backend/internal/types"
// t "github.com/BigWaffleMonster/Eventure_backend/internal/types"
// global_utils "github.com/BigWaffleMonster/Eventure_backend/internal/utils"
// "github.com/google/uuid"

type ParticipantService struct {
	repo *ParticipantRepository
}

func NewParticipantService(repo *ParticipantRepository) *ParticipantService {
	return &ParticipantService{repo: repo}
}

func (r *ParticipantService) GetParticipantsFromEvent(eventID uuid.UUID) ([]ParticipantResponse, error) {
	var participants []ParticipantResponse
	participant_raw, err := r.repo.GetParticipantsFromEvent(eventID)
	if err != nil {
		return nil, err
	}

	for _, p := range participant_raw {
		pDto := ParticipantResponse{
			ID:          p.ID,
			EventID:     p.EventID,
			DateCreated: p.DateCreated,
			User: UserResponse{
				ID:    p.User.ID,
				Login: p.User.Login,
				Email: p.User.Email,
			},
		}

		participants = append(participants, pDto)
	}

	return participants, nil
}

func (r *ParticipantService) AddParticipantToEvent(userID, eventID uuid.UUID) error {
	return nil
}

func (r *ParticipantService) RemoveParticipantFromEvent(userID, eventID uuid.UUID) error {
	return nil
}

func (r *ParticipantService) RemoveAllParticipantsFromEvent(eventID uuid.UUID) error {
	return nil
}
