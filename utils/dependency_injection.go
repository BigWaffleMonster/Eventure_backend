package utils

import "go.uber.org/fx"

func AddDI() fx.Option{
	return fx.Module(
		"Utils",
		fx.Provide(BuildServerConfig),
	)
}