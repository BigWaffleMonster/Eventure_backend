package domain_events

import (
	"fmt"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type domainEventBus struct{
	DB *gorm.DB
	Handlers []interfaces.DomainEventHandler
}

type DomainEventBusParams struct{
	fx.In

	DB *gorm.DB
	Handlers []interfaces.DomainEventHandler `group:"domainEventHandlers"`
}

func NewDomainEventBus(params DomainEventBusParams) interfaces.DomainEventBus{
	return &domainEventBus{
		DB: params.DB,
		Handlers: params.Handlers,
	}
}

func (e *domainEventBus) Publish() results.Result{
	domainEvents, result := e.GetDomainEvents()

	if result.IsFailed {
		return result
	}

	for _, domainEvent := range *domainEvents{
		result = e.publishEvent(&domainEvent)

		if result.IsFailed {
			return result.Merge(results.NewInternalError(fmt.Sprintf("failed to publish domain event: %s", domainEvent.ID.String())))
		}

		result = e.deleteFromStore(&domainEvent)

		if result.IsFailed {
			return result.Merge(results.NewInternalError(fmt.Sprintf("failed to remove domain event: %s", domainEvent.ID.String())))
		}
	}

	return results.NewResultOk()
}

func (e *domainEventBus) publishEvent(domainEventData *domain_events_base.DomainEventData) results.Result{

	for _, h := range e.Handlers {
		if h.IsTypeOf(domainEventData){
			return h.Handle(domainEventData)
		}
	}

	return results.NewResultOk()
}

func (e *domainEventBus) deleteFromStore(domainEventData *domain_events_base.DomainEventData) results.Result{
	var event domain_events_base.DomainEventData

	err := e.DB.Where("id = ?", domainEventData.ID).Delete(&event).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}


func (e *domainEventBus) GetDomainEvents() (*[]domain_events_base.DomainEventData, results.Result){
	var events []domain_events_base.DomainEventData

	result := e.DB.Find(&events)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &events, results.NewResultOk()
}
