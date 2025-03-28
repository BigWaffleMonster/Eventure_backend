package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	_ "github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/gin-gonic/gin"
)
type AuthController struct {
	Service *user.AuthService
}

func NewAuthController(service *user.AuthService) *AuthController {
	return &AuthController{Service: service}
}

// @summary register User
// @schemes
// @description register User
// @tags auth
// @accept json
// @produce json
// @param event body user.UserRegisterInput false "User"
// @success 201 {string} Successfully created!
// @failure 400 {string} string "error"
// @Router /register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var body user.UserRegisterInput

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

// @summary login User
// @schemes
// @description login User
// @tags auth
// @accept json
// @produce json
// @param event body user.UserLoginInput false "Login"
// @success 201 {string} Successfully created!
// @failure 400 {string} string "error"
// @Router /login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var body user.UserLoginInput

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
// func (c *AuthController) RefreshToken(ctx *gin.Context) {
// 	var input struct {
// 		RefreshToken string `json:"refresh_token"`
// 	}

// 	if err := ctx.ShouldBindJSON(&input); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	// Validate the refresh token
// 	claims, err := auth.ValidateRefreshToken(input.RefreshToken)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
// 		return
// 	}

// 	// Generate a new access token
// 	newAccessToken, err := auth.GenerateAccessToken(claims.Email)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating access token"})
// 		return
// 	}

// 	// Optionally generate a new refresh token
// 	newRefreshToken, err := auth.GenerateRefreshToken(claims.Email)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"access_token":  newAccessToken,
// 		"refresh_token": newRefreshToken,
// 	})
// }
