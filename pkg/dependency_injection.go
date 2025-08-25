package pkg

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"go.uber.org/fx"
)

func AddDI() fx.Option{
	return fx.Module(
		"Packages",
		fx.Provide(domain_events.NewDomainEventQueue),
		fx.Provide(domain_events.NewDomainEventBus),
		fx.Provide(domain_events.NewDomainEventStore),
	)
}