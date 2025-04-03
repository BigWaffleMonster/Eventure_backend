package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}

type authController struct {
	Service auth.AuthService
}

func NewAuthController(service auth.AuthService) AuthController {
	return &authController{Service: service}
}

func (c *authController) Register(ctx *gin.Context) {
	var body auth.RegisterInput

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

func (c *authController) Login(ctx *gin.Context) {
	var body auth.LoginInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := c.Service.Login(body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": tokens})
}

func (c *authController) RefreshToken(ctx *gin.Context) {
	var refreshInput auth.RefreshInput

	if err := ctx.ShouldBindJSON(&refreshInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	tokens, err := c.Service.RefreshToken(refreshInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": tokens})
}
