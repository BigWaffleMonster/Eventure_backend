package category

import "go.uber.org/fx"

func AddDI() fx.Option{
	return fx.Module(
		"Categories",
		fx.Provide(NewCategoryRepository),
		fx.Provide(NewCategoryService),
	)
}