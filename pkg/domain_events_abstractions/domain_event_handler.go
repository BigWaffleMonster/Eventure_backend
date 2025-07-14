package domain_events_abstractions

import "github.com/BigWaffleMonster/Eventure_backend/utils/results"

type DomainEventHandler interface{
	IsTypeOf(domainEventData *DomainEventData) bool
	Handle(domainEventData *DomainEventData) results.Result
}