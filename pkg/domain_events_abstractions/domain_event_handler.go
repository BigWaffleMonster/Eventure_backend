package domain_events_abstractions

type DomainEventHandler interface{
	IsTypeOf(domainEventData *DomainEventData) bool
	Handle(domainEventData *DomainEventData) error
}