package event

import "go.uber.org/fx"

func AddDI() fx.Option{
	return fx.Module(
		"Events",
		fx.Provide(NewEventRepository),
		fx.Provide(NewEventService),
	)
}