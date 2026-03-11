package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	c "github.com/BigWaffleMonster/Eventure_backend/internal/configs"
)

func InitRouter(config *c.Config, db *gorm.DB) *gin.Engine {
	gin.SetMode(config.Server.Mode)
	r := gin.Default()

	fmt.Print(setupCORS(config.Server.AllowedOrigins), "ETST")
	r.Use(setupCORS(config.Server.AllowedOrigins))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now(),
			"service":   "eventure-api",
		})
	})

	api := r.Group("/api/v1")
	{
		SetupAuthRoutes(api.Group("/auth"), db)
		SetupEventRoutes(api.Group("/event"), db)
		SetupParticipantsRoutes(api.Group("/participant"), db)
	}

	return r
}

func setupCORS(allowedOrigins []string) gin.HandlerFunc {
	fmt.Print(allowedOrigins)
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
