package pkg

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_bus"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_handlers"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_queue"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/migrations"
	"go.uber.org/fx"
)

func AddDI() fx.Option{
	return fx.Module(
		"Packages",
		fx.Provide(auth.NewAuthService),
		fx.Provide(migrations.InitDB),
		fx.Provide(domain_events_queue.NewDomainEventQueue),
		fx.Provide(domain_events_bus.NewDomainEventBus),
		fx.Provide(
			fx.Annotate(
				domain_events_handlers.NewEventDeletedDomainEventHandler,
				fx.ResultTags(`group:"domainEventHandlers"`),),
			),
		fx.Provide(
			fx.Annotate(
				domain_events_handlers.NewParticipantCreatedDomainEventHandler,
				fx.ResultTags(`group:"domainEventHandlers"`),),
			),
		fx.Provide(
			fx.Annotate(
				domain_events_handlers.NewUserDeletedDomainEventHandler,
				fx.ResultTags(`group:"domainEventHandlers"`),),
			),
	)
}