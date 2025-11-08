package providers

import (
	"github.com/BigWaffleMonster/Eventure_backend/config"
	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	sgloki "github.com/SergeiKhanlarov/seri-go-logger-loki"
	"github.com/SergeiKhanlarov/seri-go-logger-loki/clients"

	"go.uber.org/fx"
)

type LokiProviderParams struct {
	fx.In

    LokiClient clients.LokiClient
}

type LokiResult struct {
    fx.Out

    Provider sglogger.LoggerProvider `group:"logger_providers"`
}

func NewLokiProvider(params LokiProviderParams) LokiResult {
    level := config.GetLokiLoggingLevel()

    var minLevel sglogger.Level
    switch level {
    case "debug":
        minLevel = sglogger.LevelDebug
    case "info":
        minLevel = sglogger.LevelInfo
    case "warn":
        minLevel = sglogger.LevelWarn
    case "error":
        minLevel = sglogger.LevelError
    case "fatal":
        minLevel = sglogger.LevelFatal
    default:
        minLevel = sglogger.LevelInfo
    }

	return LokiResult{
        Provider: sgloki.NewLokiProvider(
            sgloki.ProviderConfig{
                Level: minLevel,
            },
            params.LokiClient,
        ),
    }
}

type LokiClientParams struct{
    fx.In
}

func NewLokiClient(params LokiClientParams) clients.LokiClient{
    return clients.NewLokiClient(
        clients.LokiConfig{
            LokiUrl: config.GetLokiUrl(),
            Job: "fetch",
            App: config.GetAppName(),
        },
    )
}