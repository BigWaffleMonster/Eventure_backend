package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/BigWaffleMonster/Eventure_backend/internal/api/middleware"
	configs "github.com/BigWaffleMonster/Eventure_backend/internal/configs"
	event_repo "github.com/BigWaffleMonster/Eventure_backend/internal/modules/event"
)

func SetupEventRoutes(r *gin.RouterGroup, db *gorm.DB) {
	jwtCfg := configs.InitJWTConfig()

	repo := event_repo.NewEventRepository(db)
	service := event_repo.NewEventService(repo)
	handler := event_repo.NewEventHandler(service)
	authMiddleware := middleware.NewJWTMiddleware(jwtCfg)

	r.POST("create", authMiddleware.AuthRequired(), handler.CreateEvent)

	r.GET("list", handler.GetEvents)
	r.GET(":id", handler.GetEventByID)
	r.GET("user-created/:id", handler.GetUserCreatedEvents)
	r.GET("user-participating/:id", handler.GetUserParticipatingEvents)

	r.DELETE(":id", authMiddleware.AuthRequired(), handler.RemoveEvent)

	r.PUT(":id", authMiddleware.AuthRequired(), handler.UpdateEvent)
}
