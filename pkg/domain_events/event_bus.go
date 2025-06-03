package domain_events

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"gorm.io/gorm"
)

type eventBus struct{
	DB *gorm.DB
}

func NewEventBus(db *gorm.DB) domain_events_abstractions.EventBus{
	return &eventBus{DB: db}
}

func (e *eventBus) AddToStore(domainEventData domain_events_abstractions.DomainEventData) error{
	return nil
}
func (e *eventBus) RemoveFromStore(domainEventData domain_events_abstractions.DomainEventData) error{
	return nil
}

func (e *eventBus) Publish(domainEventData domain_events_abstractions.DomainEventData) error{
	return nil
}

func (e *eventBus) GetDomainEvents() ([]domain_events_abstractions.DomainEventData, error){
	return nil, nil
}
