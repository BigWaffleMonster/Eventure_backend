package domain_events_bus

import (
	"fmt"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
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

func (e *domainEventBus) AddToStore(domainEventData *domain_events_abstractions.DomainEventData) error{
	return e.DB.Create(domainEventData).Error
}

func (e *domainEventBus) RemoveFromStore(domainEventData *domain_events_abstractions.DomainEventData) error{
	var event domain_events_abstractions.DomainEventData
	return e.DB.Where("id = ?", domainEventData.ID).Delete(&event).Error
}

func (e *domainEventBus) Publish(domainEventData *domain_events_abstractions.DomainEventData) error{

	for _, h := range e.Handlers {
		if h.IsTypeOf(domainEventData){
			return h.Handle(domainEventData)
		}
	}

	return nil
}

func (e *domainEventBus) PublishAll() error{
	domainEvents, err := e.GetDomainEvents()

	if err != nil {
		return fmt.Errorf("failed to get domain events, %e", err)
	}

	for _, domainEvent := range *domainEvents{
		err = e.Publish(&domainEvent)

		if err != nil {
			return fmt.Errorf("failed to publish domain event: %s, %e", domainEvent.ID.String(), err)
		}

		err = e.RemoveFromStore(&domainEvent)

		if err != nil {
			return fmt.Errorf("failed to remove domain event: %s, %e", domainEvent.ID.String(), err)
		}
	}

	return nil
}

func (e *domainEventBus) GetDomainEvents() (*[]domain_events_abstractions.DomainEventData, error){
	var events []domain_events_abstractions.DomainEventData

	result := e.DB.Find(&events)

	if result.Error != nil {
		return nil, result.Error
	}
	
	return &events, nil
}
