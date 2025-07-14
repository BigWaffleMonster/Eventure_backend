package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils/responses"
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
// @success 201 {object} responses.ResponseOkString
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /event [post]
func (c *EventController) Create(ctx *gin.Context) {
	var body event.EventInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to create event", []string{err.Error()}))
		return
	}

	currentUser := ctx.MustGet(auth.CurrentUserVarName).(*auth.CurrentUser)

	result := c.Service.Create(&body, currentUser)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to create event", result.Errors))
		return
	}

	ctx.JSON(http.StatusCreated, responses.NewResponseOkString("Created event success"))
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
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /event/{id} [put]
func (c *EventController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to update event", []string{err.Error()}))
		return
	}

	var body event.EventInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to update event", []string{err.Error()}))
		return
	}

	result := c.Service.Update(id, &body)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to update event", result.Errors))
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
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /event/{id} [delete]
func (c *EventController) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to delete event", []string{err.Error()}))
		return
	}

	result := c.Service.Delete(id)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to delete event", result.Errors))
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
// @success 200 {object} responses.ResponseOk[event.EventView]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /event/{id} [get]
func (c *EventController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to get event", []string{err.Error()}))
		return
	}

	eventView, result := c.Service.GetByID(id)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to get event", result.Errors))
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, responses.NewResponseOk(&eventView, "Get event success"))
}

// @Security ApiKeyAuth
// @summary Получение событий
// @schemes
// @description Получение событий
// @tags event
// @accept json
// @produce json
// @success 200 {object} responses.ResponseOk[[]event.EventView]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /event [get]
func (c *EventController) GetCollection(ctx *gin.Context) {
	eventViews, result := c.Service.GetCollection()
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to get events", result.Errors))
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, responses.NewResponseOk(&eventViews, "Get events success"))
}

// @Security ApiKeyAuth
// @summary Получение событий
// @schemes
// @description Получение событий
// @tags event
// @accept json
// @produce json
// @success 200 {object} responses.ResponseOk[[]event.EventView]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /event/private [get]
func (c *EventController) GetOwnedCollection(ctx *gin.Context) {
	currentUser := ctx.MustGet(auth.CurrentUserVarName).(*auth.CurrentUser)

	eventViews, result := c.Service.GetOwnedCollection(currentUser)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to get events", result.Errors))
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, responses.NewResponseOk(&eventViews, "Get events success"))
}