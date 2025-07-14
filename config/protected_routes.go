package config

import (
	"github.com/BigWaffleMonster/Eventure_backend/api/middlewares"
	"github.com/gin-gonic/gin"
)

func BuildProtectedRoutes(router *gin.Engine, p NewServerParams){
	protected := router.Group("/api/v1")
	protected.Use(middlewares.HandleCors(), middlewares.AuthMiddleware(p.ServerConfig))
	{
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
				participants.GET("", p.EventController.GetOwnedCollection)
			}

		}
	}
}