package domain_events_abstractions

type EventBus interface{
	AddToStore(domainEventData DomainEventData) error
	RemoveFromStore(domainEventData DomainEventData) error
	Publish(domainEventData DomainEventData) error
	GetDomainEvents() ([]DomainEventData, error)
}