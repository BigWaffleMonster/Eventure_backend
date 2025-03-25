package controller

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/models/http_models"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/service/auth_service"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/service/auth_service/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/utils"

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

	ctx.JSON(http.StatusCreated, gin.H{"data": resp})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var body http_models.UserLoginInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.Service.Login(body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": token})
}

// !TODO refresh token
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the refresh token
	claims, err := utils.ValidateRefreshToken(input.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	// Generate a new access token
	newAccessToken, err := helpers.GenerateAccessToken(claims.Email, claims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating access token"})
		return
	}

	// Optionally generate a new refresh token
	newRefreshToken, err := helpers.GenerateRefreshToken(claims.Email, claims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
