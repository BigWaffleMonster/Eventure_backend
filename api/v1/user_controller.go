package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
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
// @success 200 {object} user.UserView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /user/{id} [get]
func (c *UserController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userView, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /user/{id} [put]
func (c *UserController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body user.UserUpdateInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.Service.Update(id, &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /user/{id} [delete]
func (c *UserController) Remove(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.Service.Remove(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
