package domain_events_abstractions

type DomainEventQueue interface{
	StartQueue()
	StopQueue()
}