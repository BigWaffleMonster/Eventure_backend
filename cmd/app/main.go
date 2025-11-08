package main

import (
	v1 "github.com/BigWaffleMonster/Eventure_backend/api/v1"
	"github.com/BigWaffleMonster/Eventure_backend/cmd"
	"github.com/BigWaffleMonster/Eventure_backend/config"
	"github.com/BigWaffleMonster/Eventure_backend/internal/category"
	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/migrations"
	"github.com/BigWaffleMonster/Eventure_backend/pkg"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/logger"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/logger/providers"
	"go.uber.org/fx"
)

//	@title			Eventura app
//	@version		1.0
//	@description	Simple app to plan your celebration.

//	@contact.name   Daniil, Sergei, Alex
//	@contact.email  rachkov.work@gmail.com, Sergei.m.khanlarov@gmail.com, me@justwalsdi.ru

//	@BasePath	/api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func main() {
	config.ReadConfig()
	app := fx.New(
		event.AddDI(),
		user.AddDI(),
		participant.AddDI(),
		category.AddDI(),
		pkg.AddDI(),
		providers.AddLogger(),
		logger.AddLogger(),
		fx.Provide(migrations.InitDB),
		fx.Provide(v1.NewAuthController),
		fx.Provide(v1.NewEventController),
		fx.Provide(v1.NewCategoryController),
		fx.Provide(v1.NewUserController),
		fx.Provide(v1.NewParticipantController),
		fx.WithLogger(logger.NewFxLogger),
		fx.Invoke(cmd.NewServer),
	)

	app.Run()
}

