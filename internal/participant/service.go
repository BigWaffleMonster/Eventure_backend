package participant

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_definitions"
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/mappers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type ParticipantService interface{
    Create(ctx context.Context, data *ParticipantInput) results.Result
	ChangeState(ctx context.Context, id uuid.UUID, state *string) results.Result
	Delete(ctx context.Context, id uuid.UUID) results.Result
	GetByID(ctx context.Context, id uuid.UUID) (*ParticipantView, results.Result)
	GetCollection(ctx context.Context, eventID uuid.UUID) (*[]ParticipantView, results.Result)
	GetOwnedCollection(ctx context.Context) (*[]ParticipantView, results.Result)
}

type participantService struct {
	Uof UnitOfWork
}

func NewParticipantService(uof UnitOfWork) ParticipantService {
	return &participantService{
		Uof: uof,
	}
}

func (s *participantService) Create(ctx context.Context, data *ParticipantInput) results.Result {
	var participants *[]Participant

	participants, result := s.Uof.Repository(ctx).GetCollectionByExpression(ctx, "event_id = ?", data.EventID)

	if result.IsFailed {
		return result
	}

	currentUserID, err := helpers.GetUserID(ctx)

	if err != nil {
		return  results.NewUnauthorizedError(err.Error())
	}

	domainEventData, result := domain_events_definitions.NewUserWantsToVisitEvent(*data.EventID, currentUserID, len(*participants), *data.Status)

	if result.IsFailed {
		return result
	}

	return s.Uof.DomainEventStore(ctx).AddToStore(ctx, domainEventData)
}

func (s *participantService) ChangeState(ctx context.Context, id uuid.UUID, state *string) results.Result {
	var participant *Participant

	participant, result := s.Uof.Repository(ctx).GetByID(ctx, id)
	if result.IsFailed {
		return result
	}

	currentUserID, err := helpers.GetUserID(ctx)

	if err != nil {
		return  results.NewUnauthorizedError(err.Error())
	}

	if participant.UserID != currentUserID{
		return results.NewForbiddenError()
	}

	if state != nil {
		participant.Status = *state
	}

	return s.Uof.Repository(ctx).Update(ctx, participant)
}

func (s *participantService) Delete(ctx context.Context, id uuid.UUID) results.Result {
	participant, result := s.Uof.Repository(ctx).GetByID(ctx, id)
	if result.IsFailed {
		return result
	}

	currentUserID, err := helpers.GetUserID(ctx)

	if err != nil {
		return  results.NewUnauthorizedError(err.Error())
	}

	if participant.UserID != currentUserID{
		return results.NewForbiddenError()
	}

	return s.Uof.Repository(ctx).Delete(ctx, id)
}

func (s *participantService) GetByID(ctx context.Context, id uuid.UUID) (*ParticipantView, results.Result) {
	participant, result := s.Uof.Repository(ctx).GetByID(ctx, id)
	if result.IsFailed {
		return nil, result
	}

	currentUserID, err := helpers.GetUserID(ctx)

	if err != nil {
		return nil, results.NewUnauthorizedError(err.Error())
	}

	if participant.UserID != currentUserID{
		return nil, results.NewForbiddenError()
	}

	var participantView ParticipantView

	copier.Copy(&participantView, &participant)

	return &participantView, results.NewResultOk()
}

func (s *participantService) GetCollection(ctx context.Context, eventID uuid.UUID) (*[]ParticipantView, results.Result) {
	var participants *[]Participant

	participants, result := s.Uof.Repository(ctx).GetCollectionByExpression(ctx, "event_id = ?", eventID)
	if result.IsFailed {
		return nil, result
	}

	views := mappers.MapArray(participants, func(participant Participant) ParticipantView {
		var participantView ParticipantView
		copier.Copy(&participantView, &participant)
		return participantView
	})

	return views, results.NewResultOk()
}

func (s *participantService) GetOwnedCollection(ctx context.Context) (*[]ParticipantView, results.Result) {
	var participants *[]Participant

	currentUserID, err := helpers.GetUserID(ctx)

	if err != nil {
		return nil,  results.NewUnauthorizedError(err.Error())
	}
	
	participants, result := s.Uof.Repository(ctx).GetCollectionByExpression(ctx, "user_id=?", currentUserID)
	if result.IsFailed {
		return nil, result
	}

	views := mappers.MapArray(participants, func(participant Participant) ParticipantView {
		var participantView ParticipantView
		copier.Copy(&participantView, &participant)
		return participantView
	})

	return views, results.NewResultOk()
}