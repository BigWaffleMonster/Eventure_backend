package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
        defer cancel()
        
        c.Request = c.Request.WithContext(ctx)
        
        done := make(chan struct{})
        
        go func() {
            c.Next()
            close(done)
        }()
        
        select {
        case <-done:
            return
        case <-ctx.Done():
            if !c.Writer.Written() {
                c.AbortWithStatusJSON(408, gin.H{
                    "error": "Request timeout",
                })
            }
        }
    }
}