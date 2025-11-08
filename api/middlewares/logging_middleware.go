package middlewares

import (
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoggingMiddleware(logger sglogger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		helpers.AddToContext(c, sglogger.TraceIDKey, uuid.NewString())
		start := time.Now()
		
		path := c.Request.URL.Path
		//TODO: настраиваемые поля
		if path == "/favicon.ico" || path == "/robots.txt" {
			c.Next()
			return
		}
		
		raw := c.Request.URL.RawQuery
		
		c.Next()
		
		end := time.Now()
		latency := end.Sub(start)
		
		if raw != "" {
			path = path + "?" + raw
		}
		
		status := c.Writer.Status()
		
		logger.Info(c.Request.Context(), "HTTP Request | %3d | %13v | %15s | %-7s %s",
			status,
			latency,
			c.ClientIP(),
			c.Request.Method,
			path,
		)
	}
}
