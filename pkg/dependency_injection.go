package pkg

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/migrations"
	"go.uber.org/fx"
)

func AddDI() fx.Option{
	return fx.Module(
		"Packages",
		fx.Provide(auth.NewAuthService),
		fx.Provide(migrations.InitDB),
	)
}