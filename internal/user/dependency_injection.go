package user

import "go.uber.org/fx"

func AddDI() fx.Option{
	return fx.Module(
		"Users",
		fx.Provide(NewUserRepository),
		fx.Provide(NewUserService),
	)
}