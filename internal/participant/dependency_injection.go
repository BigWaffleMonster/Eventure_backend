package participant

import (
	"go.uber.org/fx"
)

func AddDI() fx.Option{
	return fx.Module(
		"Participants",
		fx.Provide(NewParticipantService),
		fx.Provide(NewParticipantRepository),
		fx.Provide(NewUnitOfWork),
		fx.Provide(
			fx.Annotate(
				NewEventDeletedHandler,
				fx.ResultTags(`group:"domainEventHandlers"`),),
			),
				fx.Provide(
			fx.Annotate(
				NewUserCanVisitEventHandler,
				fx.ResultTags(`group:"domainEventHandlers"`),),
			),
	)
}