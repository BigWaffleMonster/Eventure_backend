package participant

import (
	"errors"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/helpers"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type ParticipantService interface{
    Create(data *ParticipantInput, currentUser *auth.CurrentUser) error
	ChangeState(id uuid.UUID, state *string, currentUser auth.CurrentUser) error
	Remove(id uuid.UUID, currentUser auth.CurrentUser) error
	GetByID(id uuid.UUID, currentUser auth.CurrentUser) (*ParticipantView, error)
	GetCollection(eventID uuid.UUID) (*[]ParticipantView, error)
	GetOwnedCollection(currentUser *auth.CurrentUser) (*[]ParticipantView, error)
}

type participantService struct {
	Repository ParticipantRepository
}

func NewParticipantService(repository ParticipantRepository) ParticipantService {
	return &participantService{Repository: repository}
}

func (s *participantService) Create(data *ParticipantInput, currentUser *auth.CurrentUser) (error) {

	participant := Participant{
		ID: uuid.New(),
		UserID: currentUser.ID,
		EventID: *data.EventID,
		Status: *data.Status,
		Ticket: "",
	}
	
	return s.Repository.Create(&participant)
}

func (s *participantService) ChangeState(id uuid.UUID, state *string, currentUser auth.CurrentUser) (error) {
	var participant *Participant

	participant, err := s.Repository.GetByID(id)
	if err != nil {
		return err
	}

	if participant.UserID != currentUser.ID{
		return errors.New("Forbibben")//TODO: нормальное описание ошиьки и код
	}

	if state != nil {
		participant.Status = *state
	}

	return s.Repository.Update(participant)
}

func (s *participantService) Remove(id uuid.UUID, currentUser auth.CurrentUser) (error) {
	participant, err := s.Repository.GetByID(id)
	if err != nil {
		return err
	}

	if participant.UserID != currentUser.ID{
		return errors.New("Forbibben")//TODO: нормальное описание ошиьки и код
	}
	
	return s.Repository.Remove(id)
}

func (s *participantService) GetByID(id uuid.UUID, currentUser auth.CurrentUser) (*ParticipantView, error) {
	participant, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if participant.UserID != currentUser.ID{
		return nil, errors.New("Forbibben")//TODO: нормальное описание ошиьки и код
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

func (s *participantService) GetOwnedCollection(currentUser *auth.CurrentUser) (*[]ParticipantView, error) {
	var participants *[]Participant

	participants, err := s.Repository.GetOwnedCollection(currentUser.ID)
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