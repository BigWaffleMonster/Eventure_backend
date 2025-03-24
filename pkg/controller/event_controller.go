package controller

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/service"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/views"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	Service *service.EventService
}

func NewEventController(service *service.EventService) *EventController {
	return &EventController{Service: service}
}

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