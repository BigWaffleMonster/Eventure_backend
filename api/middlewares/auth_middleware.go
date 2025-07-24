package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils"
)

func AuthMiddleware(config utils.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
			c.Abort()
			return
		}

		currentUser, result := auth.ValidateAccessToken(tokenString, config)

		if result.IsFailed {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login or password"})
			c.Abort()
			return
		}

		c.Set(auth.CurrentUserVarName, currentUser)
		c.Next()
	}
}
