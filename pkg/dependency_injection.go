package pkg

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/migrations"
	"go.uber.org/fx"
)

func AddDI() fx.Option{
	return fx.Module(
		"Packages",
		//TODO: перенести
		fx.Provide(migrations.InitDB),
		fx.Provide(domain_events.NewDomainEventQueue),
		fx.Provide(domain_events.NewDomainEventBus),
		fx.Provide(domain_events.NewDomainEventStore),
	)
}