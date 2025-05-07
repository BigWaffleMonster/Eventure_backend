package config

import (
	"github.com/BigWaffleMonster/Eventure_backend/api/middlewares"
	"github.com/gin-gonic/gin"
)

func BuildPublicRoutes(router *gin.Engine, p NewServerParams){
	public := router.Group("/api/v1")
	public.Use(middlewares.HandleCors())
	{
		public.POST("/register", p.AuthController.Register)
		public.POST("/login", p.AuthController.Login)
		public.POST("/refresh", p.AuthController.RefreshToken)

		events := public.Group("/event")
		{
			events.GET("/:id", p.EventController.GetByID)
			events.GET("", p.EventController.GetCollection)
		}
		categories := public.Group("/category")
		{
			categories.GET("/:id", p.CategoryController.GetByID)
			categories.GET("", p.CategoryController.GetCollection)
		}
	}
}