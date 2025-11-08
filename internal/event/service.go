package event

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_definitions"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/mappers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type EventService interface{
    Create(ctx context.Context, data *EventInput) results.Result
	Update(ctx context.Context, id uuid.UUID, data *EventInput) results.Result
	Delete(ctx context.Context, id uuid.UUID) results.Result
	GetByID(ctx context.Context, id uuid.UUID) (*EventView, results.Result)
	GetCollection(ctx context.Context) (*[]EventView, results.Result)
	GetOwnedCollection(ctx context.Context) (*[]EventView, results.Result)
}

type eventService struct {
	UOF UnitOfWork
}

func NewEventService(UOF UnitOfWork) EventService {
	return &eventService{
		UOF: UOF,
	}
}

func (s *eventService) Create(ctx context.Context, data *EventInput) results.Result {

	if data.EndDate.Before(*data.StartDate) {
		return results.NewBadRequestError("End date is defore start date")
	}

	currentUserID, err := helpers.GetUserID(ctx)

	if err != nil {
		return  results.NewUnauthorizedError(err.Error())
	}

	event := Event{
		ID: uuid.New(),
		OwnerID: currentUserID,
		MaxQtyParticipants: *data.MaxQtyParticipants,
		Title: *data.Title,
		Description: *data.Description,
		Location: *data.Location,
		Private: *data.Private,
		StartDate: *data.StartDate,
		EndDate: *data.EndDate,
		CategoryID: *data.CategoryID,
	}
	
	return s.UOF.Repository(ctx).Create(ctx, &event)
}

func (s *eventService) Update(ctx context.Context, id uuid.UUID, data *EventInput) results.Result {
	repository := s.UOF.Repository(ctx)

	event, result := repository.GetByID(ctx, id)

	if result.IsFailed {
		return result
	}

	if data.Title != nil {
		event.Title = *data.Title
	}

	if data.Description != nil {
		event.Description = *data.Description
	}

	if data.Location != nil {
		event.Location = *data.Location
	}

	if data.Private != nil {
		event.Private = *data.Private
	}

	if data.StartDate != nil {
		event.StartDate = *data.StartDate
	}

	if data.EndDate != nil {
		event.EndDate = *data.EndDate
	}

	if data.CategoryID != nil {
		event.CategoryID = *data.CategoryID
	}

	if data.MaxQtyParticipants != nil {
		event.MaxQtyParticipants = *data.MaxQtyParticipants
	}

	return repository.Update(ctx, event)
}

func (s *eventService) Delete(ctx context.Context, id uuid.UUID) results.Result {
	return s.UOF.RunInTx(
		ctx,
		NewEventRepository,
		func(repo EventRepository, store interfaces.DomainEventStore) results.Result{
		domainEventData, result := domain_events_definitions.NewEventDeleted(id)

		if result.IsFailed {
			return result
		}

		currentUserID, err := helpers.GetUserID(ctx)

		if err != nil {
			return results.NewUnauthorizedError(err.Error())
		}

		event, result := s.UOF.Repository(ctx).GetByID(ctx, id)
		if result.IsFailed {
			return result
		}

		if event.OwnerID != currentUserID {
			return results.NewForbiddenError()
		}

		result = store.AddToStore(ctx, domainEventData)

		if result.IsFailed {
			return result
		}
		
		return repo.Delete(ctx, id)
	})
}

func (s *eventService) GetByID(ctx context.Context, id uuid.UUID) (*EventView, results.Result) {
	event, result := s.UOF.Repository(ctx).GetByID(ctx, id)
	if result.IsFailed {
		return nil, result
	}

	var eventView EventView

	copier.Copy(&eventView, &event)

	return &eventView, results.NewResultOk()
}

func (s *eventService) GetCollection(ctx context.Context) (*[]EventView, results.Result) {
	var events *[]Event

	events, result := s.UOF.Repository(ctx).GetCollection(ctx)
	if result.IsFailed {
		return nil, result
	}

	views := mappers.MapArray(events, func(event Event) EventView {
		var eventView EventView
		copier.Copy(&eventView, &event)
		return eventView
	})

	return views, results.NewResultOk()
}

func (s *eventService) GetOwnedCollection(ctx context.Context) (*[]EventView, results.Result) {
	var events *[]Event

	currentUserID, err := helpers.GetUserID(ctx)

	if err != nil {
		return nil, results.NewUnauthorizedError(err.Error())
	}

	events, result := s.UOF.Repository(ctx).GetCollectionByExpression(ctx, "owner_id = ?", currentUserID)
	if result.IsFailed {
		return nil, result
	}

	views := mappers.MapArray(events, func(event Event) EventView {
		var eventView EventView
		copier.Copy(&eventView, &event)
		return eventView
	})

	return views, results.NewResultOk()
}