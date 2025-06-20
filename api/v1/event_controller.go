package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventController struct {
	Service event.EventService
}

func NewEventController(service event.EventService) *EventController {
	return &EventController{Service: service}
}

// @Security ApiKeyAuth
// @summary Создание события
// @schemes
// @description Создание события
// @tags event
// @accept json
// @produce json
// @param event body event.EventInput false "Событие"
// @success 204 
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event [post]
func (c *EventController) Create(ctx *gin.Context) {
	var body event.EventInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := ctx.MustGet(auth.CurrentUserVarName).(*auth.CurrentUser)

	err := c.Service.Create(&body, currentUser)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @Security ApiKeyAuth
// @summary Обновление события
// @schemes
// @description Обновление события
// @tags event
// @accept json
// @produce json
// @param id path string true "Event ID"
// @param event body event.EventInput false "Событие"
// @success 204 
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event/{id} [put]
func (c *EventController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body event.EventInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.Service.Update(id, &body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @Security ApiKeyAuth
// @summary Удаление события
// @schemes
// @description Удаление события
// @tags event
// @accept json
// @produce json
// @param id path string true "Идентификатор события"
// @success 204 
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event/{id} [delete]
func (c *EventController) Remove(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.Service.Remove(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @Security ApiKeyAuth
// @summary Получение события
// @schemes
// @description Получение события
// @tags event
// @accept json
// @produce json
// @param id path string true "Идентификатор события"
// @success 200 {object} event.EventView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event/{id} [get]
func (c *EventController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventView, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, eventView)
}

// @Security ApiKeyAuth
// @summary Получение событий
// @schemes
// @description Получение событий
// @tags event
// @accept json
// @produce json
// @success 200 {object} []event.EventView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event [get]
func (c *EventController) GetCollection(ctx *gin.Context) {
	eventViews, err := c.Service.GetCollection()
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, eventViews)
}

// @Security ApiKeyAuth
// @summary Получение событий
// @schemes
// @description Получение событий
// @tags event
// @accept json
// @produce json
// @success 200 {object} []event.EventView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event/private [get]
func (c *EventController) GetOwnedCollection(ctx *gin.Context) {
	currentUser := ctx.MustGet(auth.CurrentUserVarName).(*auth.CurrentUser)

	eventViews, err := c.Service.GetOwnedCollection(currentUser)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, eventViews)
}