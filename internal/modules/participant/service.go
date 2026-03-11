package participant

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/utils"

	"github.com/google/uuid"
)

// "net/http"
// "time"

// "github.com/BigWaffleMonster/Eventure_backend/internal/types"
// t "github.com/BigWaffleMonster/Eventure_backend/internal/types"
// "github.com/google/uuid"

type ParticipantService struct {
	repo *ParticipantRepository
}

func NewParticipantService(repo *ParticipantRepository) *ParticipantService {
	return &ParticipantService{repo: repo}
}

func (s *ParticipantService) GetParticipantsFromEvent(eventID uuid.UUID) ([]ParticipantResponse, error) {
	var participants []ParticipantResponse
	participant_raw, err := s.repo.GetParticipantsFromEvent(eventID)
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

func (s *ParticipantService) AddParticipantToEvent(userID, eventID uuid.UUID) error {
	err := s.repo.CheckUserExistence(userID)
	if err != nil {
		return err
	}

	err = s.repo.CheckIfUserParticipant(userID, eventID)
	if err != nil {
		return err
	}

	cap, maxCap, err := s.repo.GetEventCapacity(eventID)
	if err != nil {
		return err
	}

	if maxCap != nil {
		if int64(*cap) >= int64(*maxCap) {
			return utils.NewAppErrorWithErr(
				http.StatusBadRequest,
				"Event has reached maximum capacity",
				nil,
			)
		}
	}

	err = s.repo.AddParticipantToEvent(userID, eventID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ParticipantService) RemoveParticipantFromEvent(userID, eventID uuid.UUID) error {
	err := s.repo.RemoveParticipantFromEvent(userID, eventID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ParticipantService) RemoveAllParticipantsFromEvent(eventID uuid.UUID) error {
	err := s.repo.RemoveAllParticipantsFromEvent(eventID)
	if err != nil {
		return err
	}

	return nil
}
