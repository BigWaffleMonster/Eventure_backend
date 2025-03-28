package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventController struct {
	Service *event.EventService
}

func NewEventController(service *event.EventService) *EventController {
	return &EventController{Service: service}
}

// @summary create Event
// @schemes
// @description create Event
// @tags event
// @accept json
// @produce json
// @param event body event.EventInput false "Event"
// @success 204 
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event [post]
func (controller *EventController) Create(ctx *gin.Context) {
	var body event.EventInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.Service.Create(body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @summary update Event
// @schemes
// @description update Event
// @tags event
// @accept json
// @produce json
// @param id path string true "Event ID"
// @param event body event.EventInput false "Event"
// @success 204 
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event/{id} [put]
func (controller *EventController) Update(ctx *gin.Context) {
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

	err = controller.Service.Update(id, body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @summary delete Event
// @schemes
// @description delete Event
// @tags event
// @accept json
// @produce json
// @param id path string true "Event ID"
// @success 204 
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event/{id} [delete]
func (controller *EventController) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @summary get Event
// @schemes
// @description get Event
// @tags event
// @accept json
// @produce json
// @param id path string true "Event ID"
// @success 200 {object} event.EventView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event/{id} [get]
func (controller *EventController) GetById(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventView, err := controller.Service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, eventView)
}

// @summary get Events
// @schemes
// @description get Events
// @tags event
// @accept json
// @produce json
// @success 200 {object} []event.EventView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /event [get]
func (controller *EventController) GetCollection(ctx *gin.Context) {
	eventViews, err := controller.Service.GetCollection()
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, eventViews)
}