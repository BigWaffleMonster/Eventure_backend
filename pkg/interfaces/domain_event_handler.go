package interfaces

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
)

type DomainEventHandler interface{
	IsTypeOf(ctx context.Context, domainEventData *domain_events_base.DomainEventData) bool
	Handle(ctx context.Context, domainEventData *domain_events_base.DomainEventData) results.Result
}