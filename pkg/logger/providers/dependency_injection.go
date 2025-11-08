package providers

import (
	"go.uber.org/fx"
)

func AddLogger() fx.Option {
	return fx.Module("providers",
		fx.Provide(
			NewLokiProvider,
			NewConsoleProvider,
			NewConsoleFormatter,
			NewLokiClient,
		),
	)
}