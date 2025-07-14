package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/utils/responses"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	Service user.UserService
}

func NewUserController(service user.UserService) *UserController {
	return &UserController{Service: service}
}

//TODO: а зачем тут айди? получаем айжи не с фронта, а изнутри контекста запроса
// @Security ApiKeyAuth
// @summary Получение пользователя
// @schemes
// @description Получение пользователя
// @tags user
// @accept json
// @produce json
// @param id path string true "Идентиикатор пользователя"
// @success 200 {object} responses.ResponseOk[user.UserView]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /user/{id} [get]
func (c *UserController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to get user", []string{err.Error()}))
		return
	}

	userView, result := c.Service.GetByID(id)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to get user", result.Errors))
		return
	}

	ctx.JSON(http.StatusOK, userView)
}

// @Security ApiKeyAuth
// @summary Обновление пользователя
// @schemes
// @description Обновление пользователя
// @tags user
// @accept json
// @produce json
// @param id path string true "Идентиикатор пользователя"
// @param event body user.UserUpdateInput false "Данные о пользователе"
// @success 204 
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /user/{id} [put]
func (c *UserController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to update user", []string{err.Error()}))
		return
	}

	var body user.UserUpdateInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := c.Service.Update(id, &body)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to update user", result.Errors))
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @Security ApiKeyAuth
// @summary Удаление пользователя
// @schemes
// @description Удаление пользователя
// @tags user
// @accept json
// @produce json
// @param id path string true "Идентиикатор пользователя"
// @success 204 
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /user/{id} [delete]
func (c *UserController) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to delete user", []string{err.Error()}))
		return
	}

	result := c.Service.Delete(id)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to delete user", result.Errors))
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
