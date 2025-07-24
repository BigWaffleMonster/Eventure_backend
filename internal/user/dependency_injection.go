package user

import "go.uber.org/fx"

func AddDI() fx.Option{
	return fx.Module(
		"Users",
		fx.Provide(NewUserService),
		fx.Provide(NewUnitOfWork),
		fx.Provide(NewUserRepository),
		fx.Provide(NewAuthService),
	)
}