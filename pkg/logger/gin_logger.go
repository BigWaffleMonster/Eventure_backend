package logger

import (
	"context"
	"io"
	"strings"

	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
)

type ginToMyLogger struct {
    logger sglogger.Logger
	context context.Context
}

func NewGinLogger(Logger sglogger.Logger) io.Writer{
	return &ginToMyLogger{
		logger: Logger,
		context: context.Background(),
	}
}

func (g *ginToMyLogger) Write(p []byte) (n int, err error) {
	    line := strings.TrimSpace(string(p))
    if line == "" {
        return len(p), nil
    }
    
    // Определяем уровень по содержимому строки Gin
    switch {
    case strings.Contains(line, "[GIN-debug]"):
        msg := strings.TrimPrefix(line, "[GIN-debug] ")
        g.logger.Debug(g.context, msg)
        
    case strings.Contains(line, "[WARNING]"):
        msg := strings.TrimPrefix(line, "[WARNING] ")
        g.logger.Warning(g.context, msg)
        
    case strings.Contains(line, "[ERROR]"):
        msg := strings.TrimPrefix(line, "[ERROR] ")
        g.logger.Error(g.context, msg)
        
    case strings.Contains(line, "[GIN]"):
        msg := strings.TrimPrefix(line, "[GIN] ")
        g.logger.Info(g.context, msg)
		g.logger.InfoWithFields(g.context, sglogger.Fields{"gin": true}, line)
        
    default:
        g.logger.Info(g.context, line)
    }
    
    return len(p), nil
    // line := strings.TrimSpace(string(p))
    // if line != "" {
	// 	g.logger.Info(g.context, line)
    //     //g.logger.InfoWithFields(g.context, sglogger.Fields{"gin": true}, line)
    // }
    // return len(p), nil
}