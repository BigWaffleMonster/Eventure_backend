package interfaces

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
)

type DomainEventHandler interface{
	IsTypeOf(domainEventData *domain_events_base.DomainEventData) bool
	Handle(domainEventData *domain_events_base.DomainEventData) results.Result
}