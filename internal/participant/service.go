package participant

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/helpers"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type ParticipantService interface{
    Create(data *ParticipantInput) error
	ChangeState(id uuid.UUID, state *string) error
	Remove(id uuid.UUID) error
	GetByID(id uuid.UUID) (*ParticipantView, error)
	GetCollection(eventID uuid.UUID) (*[]ParticipantView, error)
}

type participantService struct {
	Repository ParticipantRepository
}

func NewParticipantService(repository ParticipantRepository) ParticipantService {
	return &participantService{Repository: repository}
}

func (s *participantService) Create(data *ParticipantInput) (error) {

	participant := Participant{
		ID: uuid.New(),
		UserID: uuid.New(),//TODO: add un middleware and get user from it
		EventID: *data.EventID,
		Status: *data.Status,
		Ticket: "",
	}
	
	return s.Repository.Create(&participant)
}

func (s *participantService) ChangeState(id uuid.UUID, state *string) (error) {
	var participant *Participant

	participant, err := s.Repository.GetByID(id)
	if err != nil {
		return err
	}

	if state != nil {
		participant.Status = *state
	}

	return s.Repository.Update(participant)
}

func (s *participantService) Remove(id uuid.UUID) (error) {
	return s.Repository.Remove(id)
}

func (s *participantService) GetByID(id uuid.UUID) (*ParticipantView, error) {
	participant, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	var participantView ParticipantView

	copier.Copy(&participantView, &participant)

	return &participantView, nil
}

func (s *participantService) GetCollection(eventID uuid.UUID) (*[]ParticipantView, error) {
	var participants *[]Participant

	participants, err := s.Repository.GetCollection(eventID)
	if err != nil {
		return nil, err
	}

	views := helpers.MapArray(participants, func(participant Participant) ParticipantView {
		var participantView ParticipantView
		copier.Copy(&participantView, &participant)
		return participantView
	})

	return views, nil
}