package cmd

import (
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/api/middlewares"
	"github.com/BigWaffleMonster/Eventure_backend/config"
	"github.com/gin-gonic/gin"
)
func BuildRoutes(router *gin.Engine, p NewServerParams) {
    api := router.Group("/api/v1")
    api.Use(
        middlewares.LoggingMiddleware(p.Logger),
		middlewares.CorsMiddleware(), 
		middlewares.RequestInfoMiddleware(), 
		middlewares.TimeoutMiddleware(time.Duration(config.GetRequestTimeout()) * time.Second))
    
    buildPublicRoutes(api, p)
    buildProtectedRoutes(api, p)
}

func buildPublicRoutes(api *gin.RouterGroup, p NewServerParams) {
    api.POST("/logout", p.AuthController.Logout)
    api.POST("/register", p.AuthController.Register)
    api.POST("/login", p.AuthController.Login)
    api.POST("/refresh", p.AuthController.RefreshToken)

    events := api.Group("/event")
    {
        events.GET("/:id", p.EventController.GetByID)
        events.GET("", p.EventController.GetCollection)
    }
    categories := api.Group("/category")
    {
        categories.GET("/:id", p.CategoryController.GetByID)
        categories.GET("", p.CategoryController.GetCollection)
    }
}

func buildProtectedRoutes(api *gin.RouterGroup, p NewServerParams) {
    protected := api.Group("")
    protected.Use(middlewares.AuthMiddleware())
    {
        user := protected.Group("/user")
        {
            user.GET("/:id", p.UserController.GetByID)
            user.PUT("/:id", p.UserController.Update)
            user.DELETE("/:id", p.UserController.Delete)
        }
        events := protected.Group("/event")
        {
            events.POST("", p.EventController.Create)
            events.PUT("/:id", p.EventController.Update)
            events.DELETE("/:id", p.EventController.Delete)
            events.GET("/private", p.EventController.GetOwnedCollection)
        }
        participants := protected.Group("/participant")
        {
            participants.POST("", p.ParticipantController.Create)
            participants.PUT("/:id/state", p.ParticipantController.ChangeState)
            participants.DELETE("/:id", p.ParticipantController.Delete)
            participants.GET("/:id", p.ParticipantController.GetByID)
            participants.GET("/event/:eventId", p.ParticipantController.GetCollection)
            participants.GET("", p.ParticipantController.GetOwnedCollection)
        }
    }
}