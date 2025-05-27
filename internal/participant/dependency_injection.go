package participant

import "go.uber.org/fx"

func AddDI() fx.Option{
	return fx.Module(
		"Participants",
		fx.Provide(NewParticipantRepository),
		fx.Provide(NewParticipantService),
	)
}