package domain_events_abstractions

type DomainEventBus interface{
	AddToStore(domainEventData *DomainEventData) error
	RemoveFromStore(domainEventData *DomainEventData) error
	Publish(domainEventData *DomainEventData) error
	GetDomainEvents() (*[]DomainEventData, error)
	PublishAll() error
}