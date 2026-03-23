package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	configs "github.com/BigWaffleMonster/Eventure_backend/internal/configs"
	user_repo "github.com/BigWaffleMonster/Eventure_backend/internal/modules/auth"
)

func SetupAuthRoutes(r *gin.RouterGroup, db *gorm.DB) {
	jwtCfg := configs.InitJWTConfig()

	repo := user_repo.NewAuthRepository(db)
	service := user_repo.NewAuthService(repo, jwtCfg)
	handler := user_repo.NewAuthHandler(service)

	r.POST("sign-up", handler.CreateUser)         // imp
	r.POST("sign-in", handler.Login)              // imp
	r.POST("logout", handler.Logout)              // not-imp
	r.POST("refresh-token", handler.RefreshToken) // imp
}
