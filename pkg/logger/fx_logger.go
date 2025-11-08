package logger

import (
	"context"

	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	"go.uber.org/fx/fxevent"
)

type FxLogger struct {
    logger sglogger.Logger
}

func NewFxLogger(logger sglogger.Logger) fxevent.Logger {
    return &FxLogger{logger: logger}
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
    ctx := context.Background()
    switch e := event.(type) {
    case *fxevent.OnStartExecuting:
        l.logger.Debug(ctx,
            "OnStart hook executing",
            "callee", e.FunctionName,
            "caller", e.CallerName)
    case *fxevent.OnStartExecuted:
        if e.Err != nil {
            l.logger.Error(ctx,
                "OnStart hook failed",
                "callee", e.FunctionName,
                "caller", e.CallerName,
                "error", e.Err)
        } else {
            l.logger.Debug(ctx,
                "OnStart hook executed",
                "callee", e.FunctionName,
                "caller", e.CallerName,
                "runtime", e.Runtime)
        }
    case *fxevent.OnStopExecuting:
        l.logger.Debug(ctx,
            "OnStop hook executing",
            "callee", e.FunctionName,
            "caller", e.CallerName)
    case *fxevent.OnStopExecuted:
        if e.Err != nil {
            l.logger.Error(ctx,
                "OnStop hook failed",
                "callee", e.FunctionName,
                "caller", e.CallerName,
                "error", e.Err)
        } else {
            l.logger.Debug(ctx,
                "OnStop hook executed",
                "callee", e.FunctionName,
                "caller", e.CallerName,
                "runtime", e.Runtime)
        }
    case *fxevent.Supplied:
        if e.Err != nil {
            l.logger.Error(ctx,
                "Error encountered while applying options",
                "type", e.TypeName,
                "error", e.Err)
        } else {
            l.logger.Debug(ctx, "Supplied", "type", e.TypeName)
        }
    case *fxevent.Provided:
        for _, rtype := range e.OutputTypeNames {
            l.logger.Debug(ctx,
                "Provided",
                "constructor", e.ConstructorName,
                "type", rtype)
        }
        if e.Err != nil {
            l.logger.Error(ctx,
                "Error encountered while applying options",
                "error", e.Err)
        }
    case *fxevent.Replaced:
        for _, rtype := range e.OutputTypeNames {
            l.logger.Debug(ctx,
                "Replaced",
                "type", rtype)
        }
        if e.Err != nil {
            l.logger.Error(ctx,
                "Error encountered while replacing",
                "error", e.Err)
        }
    case *fxevent.Decorated:
        for _, rtype := range e.OutputTypeNames {
            l.logger.Debug(ctx,
                "Decorated",
                "decorator", e.DecoratorName,
                "type", rtype)
        }
        if e.Err != nil {
            l.logger.Error(ctx,
                "Error encountered while decorating",
                "error", e.Err)
        }
    case *fxevent.Invoked:
        if e.Err != nil {
            l.logger.Error(ctx,
                "Invoke failed",
                "function", e.FunctionName,
                "error", e.Err)
        } else {
            l.logger.Debug(ctx,"Invoked", "function", e.FunctionName)
        }
    case *fxevent.Stopping:
        l.logger.Info(ctx, "Received signal", "signal", e.Signal)
    case *fxevent.Stopped:
        if e.Err != nil {
            l.logger.Error(ctx,"Stop failed", "error", e.Err)
        } else {
            l.logger.Debug(ctx, "Stopped")
        }
    case *fxevent.RollingBack:
        l.logger.Error(ctx, "Start failed, rolling back", "error", e.StartErr)
    case *fxevent.RolledBack:
        if e.Err != nil {
            l.logger.Error(ctx, "Rollback failed", "error", e.Err)
        } else {
            l.logger.Debug(ctx, "Rolled back")
        }
    case *fxevent.Started:
        if e.Err != nil {
            l.logger.Error(ctx,"Start failed", "error", e.Err)
        } else {
            l.logger.Debug(ctx, "Started")
        }
    case *fxevent.LoggerInitialized:
        if e.Err != nil {
            l.logger.Error(ctx,"Custom logger initialization failed", "error", e.Err)
        } else {
            l.logger.Debug(ctx, "Initialized custom fx event logger")
        }
    }
}