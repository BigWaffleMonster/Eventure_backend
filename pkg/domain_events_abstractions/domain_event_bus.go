package domain_events_abstractions

import "github.com/BigWaffleMonster/Eventure_backend/utils/results"

type DomainEventBus interface{
	AddToStore(domainEventData *DomainEventData) results.Result
	DeleteFromStore(domainEventData *DomainEventData) results.Result
	Publish(domainEventData *DomainEventData) results.Result
	GetDomainEvents() (*[]DomainEventData, results.Result)
	PublishAll() results.Result
}