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

// @summary Зарегестрирвоаться
// @schemes
// @description Зарегестрирвоаться
// @tags auth
// @accept json
// @produce json
// @param register body auth.RegisterInput false "Данные о пользоавтеле для регистрации"
// @success 201 {string} string "created"
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /register [post]
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

// @summary Войти в систему
// @schemes
// @description Войти в систему
// @tags auth
// @accept json
// @produce json
// @param register body auth.LoginInput false "Данные логина"
// @success 200 {string} string "created"
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /login [post]
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

	ctx.JSON(http.StatusOK, gin.H{"data": tokens})
}

// @summary Обновить токен
// @schemes
// @description Обновить токен
// @tags auth
// @accept json
// @produce json
// @param register body auth.RefreshInput false "Обновить токен"
// @success 200 {string} string "created"
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /refresh [post]
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
