package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils/responses"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service user.AuthService
}

func NewAuthController(service user.AuthService) *AuthController {
	return &AuthController{Service: service}
}

// @summary Регистрация нового пользователя
// @schemes
// @description Регистрация нового пользователя
// @tags auth
// @accept json
// @produce json
// @param register body auth.RegisterInput false "Данные о пользоавтеле для регистрации"
// @success 201 {object} responses.ResponseOkString
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var body auth.RegisterInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to register", []string{err.Error()}))
		return
	}

	result := c.Service.Register(body)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to register", result.Errors))
		return
	}

	ctx.JSON(http.StatusCreated, responses.NewResponseOkString("Regictered success"))
}

// @summary Войти в систему
// @schemes
// @description Войти в систему
// @tags auth
// @accept json
// @produce json
// @param register body auth.LoginInput false "Данные логина"
// @success 200 {object} responses.ResponseOk[[]string]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var body auth.LoginInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to login", []string{err.Error()}))
		return
	}

	tokens, result := c.Service.Login(body)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to login", result.Errors))
		return
	}

	ctx.JSON(http.StatusOK, responses.NewResponseOk(&tokens, "Login success"))
}

// @summary Обновить токен
// @schemes
// @description Обновить токен
// @tags auth
// @accept json
// @produce json
// @param register body auth.RefreshInput false "Обновить токен"
// @success 200 {object} responses.ResponseOk[[]string]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /refresh [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var refreshInput auth.RefreshInput

	if err := ctx.ShouldBindJSON(&refreshInput); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to refresh token", []string{err.Error()}))
		return
	}

	tokens, result := c.Service.RefreshToken(refreshInput)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to refresh token", result.Errors))
		return
	}

	ctx.JSON(http.StatusOK, responses.NewResponseOk(&tokens, "Refresh token success"))
}
