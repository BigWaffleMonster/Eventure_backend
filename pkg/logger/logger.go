package logger

import (
	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	"go.uber.org/fx"
)

type LoggerParams struct {
	fx.In

    Providers []sglogger.LoggerProvider `group:"logger_providers"`
}

func NewLogger(params LoggerParams) sglogger.Logger {
    return sglogger.NewLogger(
        sglogger.LoggerConfig{},
        sglogger.NewFieldsHandler(),
        params.Providers...
    )
}