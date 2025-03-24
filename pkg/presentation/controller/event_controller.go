package controller

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/application/services"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/application/views"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	Service *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{Service: service}
}

// @summary create Event
// @schemes
// @description create Event
// @tags example
// @accept json
// @produce json
// @param event body views.EventInfo false "Event"
// @success 201 {string} Successfully created!
// @failure 400 {string} string "error"
// @router /event [post]
func (c *EventController) Create(ctx *gin.Context) {
	var body views.EventInfo

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.Service.Create(body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": resp})
}