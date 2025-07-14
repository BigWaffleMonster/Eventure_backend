package domain_events_bus

import (
	"fmt"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type domainEventBus struct{
	DB *gorm.DB
	Handlers []domain_events_abstractions.DomainEventHandler
}

type DomainEventBusParams struct{
	fx.In

	DB *gorm.DB
	Handlers []domain_events_abstractions.DomainEventHandler `group:"domainEventHandlers"`
}

func NewDomainEventBus(params DomainEventBusParams) domain_events_abstractions.DomainEventBus{
	return &domainEventBus{
		DB: params.DB,
		Handlers: params.Handlers,
	}
}

func (e *domainEventBus) AddToStore(domainEventData *domain_events_abstractions.DomainEventData) results.Result{
	result := e.DB.Create(domainEventData).Error

	if result != nil {
		return results.NewInternalError(result.Error())
	}

	return results.NewResultOk()
}

func (e *domainEventBus) DeleteFromStore(domainEventData *domain_events_abstractions.DomainEventData) results.Result{
	var event domain_events_abstractions.DomainEventData

	err := e.DB.Where("id = ?", domainEventData.ID).Delete(&event).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (e *domainEventBus) Publish(domainEventData *domain_events_abstractions.DomainEventData) results.Result{

	for _, h := range e.Handlers {
		if h.IsTypeOf(domainEventData){
			return h.Handle(domainEventData)
		}
	}

	return results.NewResultOk()
}

func (e *domainEventBus) PublishAll() results.Result{
	domainEvents, result := e.GetDomainEvents()

	if result.IsFailed {
		return result
	}

	for _, domainEvent := range *domainEvents{
		result = e.Publish(&domainEvent)

		if result.IsFailed {
			return result.Merge(results.NewInternalError(fmt.Sprintf("failed to publish domain event: %s", domainEvent.ID.String())))
		}

		result = e.DeleteFromStore(&domainEvent)

		if result.IsFailed {
			return result.Merge(results.NewInternalError(fmt.Sprintf("failed to remove domain event: %s", domainEvent.ID.String())))
		}
	}

	return results.NewResultOk()
}

func (e *domainEventBus) GetDomainEvents() (*[]domain_events_abstractions.DomainEventData, results.Result){
	var events []domain_events_abstractions.DomainEventData

	result := e.DB.Find(&events)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &events, results.NewResultOk()
}
