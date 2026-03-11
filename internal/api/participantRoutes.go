package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/BigWaffleMonster/Eventure_backend/internal/api/middleware"
	configs "github.com/BigWaffleMonster/Eventure_backend/internal/configs"
	participant_repo "github.com/BigWaffleMonster/Eventure_backend/internal/modules/participant"
)

// TODO Test routes fot bugs. Base functions are working fine
// Check for that current user only can operate with himsels (expect owners and admins)
func SetupParticipantsRoutes(r *gin.RouterGroup, db *gorm.DB) {
	jwtCfg := configs.InitJWTConfig()

	repo := participant_repo.NewParticipantRepository(db)
	service := participant_repo.NewParticipantService(repo)
	handler := participant_repo.NewParticipantHandler(service)
	authMiddleware := middleware.NewJWTMiddleware(jwtCfg)

	r.GET(":id", handler.GetParticipantsFromEvent)

	// TODO owner cant be a participant to his events
	// And only current user can add himself
	r.POST("add", authMiddleware.AuthRequired(), handler.AddParticipantToEvent)

	// TODO only user can remove himeself of owner of the event or admin
	r.DELETE("", authMiddleware.AuthRequired(), handler.RemoveParticipantFromEvent)
	// TODO only owner of event or admin can remove
	r.DELETE("remove_all/:id", authMiddleware.AuthRequired(), handler.RemoveAllParticipantsFromEvent)
}
