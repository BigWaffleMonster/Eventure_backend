package event

import "go.uber.org/fx"

func AddDI() fx.Option{
	return fx.Module(
		"Events",
		fx.Provide(NewEventService),
		fx.Provide(NewUnitOfWork),
		fx.Provide(
			fx.Annotate(
				NewUserDeletedHandler,
				fx.ResultTags(`group:"domainEventHandlers"`),),
			),
		fx.Provide(
			fx.Annotate(
				NewUserWantsToVisitEventHandler,
				fx.ResultTags(`group:"domainEventHandlers"`),),
			),
	)
}