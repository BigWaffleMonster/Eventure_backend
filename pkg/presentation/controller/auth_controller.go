package controller

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/application/services"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/application/views"
	"github.com/gin-gonic/gin"
)

type AuthController struct{
	Service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{Service: service}
}


// @summary register User
// @schemes
// @description register User
// @tags example
// @accept json
// @produce json
// @param event body views.UserInfo false "User"
// @success 201 {string} Successfully created!
// @failure 400 {string} string "error"
// @Router /register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var body views.UserInfo

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

// @summary login User
// @schemes
// @description login User
// @tags example
// @accept json
// @produce json
// @param event body views.LoginInfo false "Login"
// @success 201 {string} Successfully created!
// @failure 400 {string} string "error"
// @Router /register [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var body views.LoginInfo

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