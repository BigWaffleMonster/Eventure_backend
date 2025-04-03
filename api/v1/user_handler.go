package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	GetUserByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Remove(ctx *gin.Context)
}

type userController struct {
	Service user.UserService
}

func NewUserController(service user.UserService) UserController {
	return &userController{Service: service}
}

func (c *userController) GetUserByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (c *userController) Update(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *userController) Remove(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, gin.H{})
}
