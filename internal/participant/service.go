package participant

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_definitions"
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type ParticipantService interface{
    Create(data *ParticipantInput, currentUser *auth.CurrentUser) results.Result
	ChangeState(id uuid.UUID, state *string, currentUser auth.CurrentUser) results.Result
	Delete(id uuid.UUID, currentUser auth.CurrentUser) results.Result
	GetByID(id uuid.UUID, currentUser auth.CurrentUser) (*ParticipantView, results.Result)
	GetCollection(eventID uuid.UUID) (*[]ParticipantView, results.Result)
	GetOwnedCollection(currentUser *auth.CurrentUser) (*[]ParticipantView, results.Result)
}

type participantService struct {
	Uof UnitOfWork
}

func NewParticipantService(uof UnitOfWork) ParticipantService {
	return &participantService{
		Uof: uof,
	}
}

func (s *participantService) Create(data *ParticipantInput, currentUser *auth.CurrentUser) results.Result {
	var participants *[]Participant

	participants, result := s.Uof.Repository().GetCollectionByExpression("event_id = ?", data.EventID)

	if result.IsFailed {
		return result
	}

	domainEventData, result := domain_events_definitions.NewUserWantsToVisitEvent(*data.EventID, currentUser.ID, len(*participants), *data.Status)

	if result.IsFailed {
		return result
	}

	return s.Uof.DomainEventStore().AddToStore(domainEventData)
}

func (s *participantService) ChangeState(id uuid.UUID, state *string, currentUser auth.CurrentUser) results.Result {
	var participant *Participant

	participant, result := s.Uof.Repository().GetByID(id)
	if result.IsFailed {
		return result
	}

	if participant.UserID != currentUser.ID{
		return results.NewForbiddenError()
	}

	if state != nil {
		participant.Status = *state
	}

	return s.Uof.Repository().Update(participant)
}

func (s *participantService) Delete(id uuid.UUID, currentUser auth.CurrentUser) results.Result {
	participant, result := s.Uof.Repository().GetByID(id)
	if result.IsFailed {
		return result
	}

	if participant.UserID != currentUser.ID{
		return results.NewForbiddenError()
	}

	return s.Uof.Repository().Delete(id)
}

func (s *participantService) GetByID(id uuid.UUID, currentUser auth.CurrentUser) (*ParticipantView, results.Result) {
	participant, result := s.Uof.Repository().GetByID(id)
	if result.IsFailed {
		return nil, result
	}

	if participant.UserID != currentUser.ID{
		return nil, results.NewForbiddenError()
	}

	var participantView ParticipantView

	copier.Copy(&participantView, &participant)

	return &participantView, results.NewResultOk()
}

func (s *participantService) GetCollection(eventID uuid.UUID) (*[]ParticipantView, results.Result) {
	var participants *[]Participant

	participants, result := s.Uof.Repository().GetCollectionByExpression("event_id = ?", eventID)
	if result.IsFailed {
		return nil, result
	}

	views := helpers.MapArray(participants, func(participant Participant) ParticipantView {
		var participantView ParticipantView
		copier.Copy(&participantView, &participant)
		return participantView
	})

	return views, results.NewResultOk()
}

func (s *participantService) GetOwnedCollection(currentUser *auth.CurrentUser) (*[]ParticipantView, results.Result) {
	var participants *[]Participant
	
	participants, result := s.Uof.Repository().GetCollectionByExpression("user_id=?", currentUser.ID)
	if result.IsFailed {
		return nil, result
	}

	views := helpers.MapArray(participants, func(participant Participant) ParticipantView {
		var participantView ParticipantView
		copier.Copy(&participantView, &participant)
		return participantView
	})

	return views, results.NewResultOk()
}