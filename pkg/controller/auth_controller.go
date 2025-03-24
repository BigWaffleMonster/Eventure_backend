package controller

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/models/http_models"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/service/auth_service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *auth_service.AuthService
}

func NewAuthController(service *auth_service.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var body http_models.UserRegisterInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.Service.Register(body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": resp})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var body http_models.UserLoginInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.Service.Login(body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": resp})
}
