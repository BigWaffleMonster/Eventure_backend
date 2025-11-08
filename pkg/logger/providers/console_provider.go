package providers

import (
	"github.com/BigWaffleMonster/Eventure_backend/config"
	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	sgconsole "github.com/SergeiKhanlarov/seri-go-logger-console"
	"go.uber.org/fx"
)

type ConsoleProviderParams struct {
	fx.In

	Formatter sgconsole.ConsoleFormatter
    
}

type ConsoleProviderResult struct {
    fx.Out

    Provider sglogger.LoggerProvider `group:"logger_providers"`
}


func NewConsoleProvider(params ConsoleProviderParams) ConsoleProviderResult {    
    level := config.GetConsoleLoggingLevel()

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

	return ConsoleProviderResult{
        Provider: sgconsole.NewConsoleProvider(
            sgconsole.ProviderConfig{
                Level: minLevel,
            },
            params.Formatter,
        ),
    }
}

type ConsoleFormatterParams struct{
    fx.In
}

func NewConsoleFormatter(params ConsoleFormatterParams) sgconsole.ConsoleFormatter{
    return sgconsole.NewConsoleFormatter()
}