package event

import (
	"net/http"

	utils "github.com/BigWaffleMonster/Eventure_backend/internal/modules/event/utils"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service *EventService
}

func NewEventHandler(service *EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req CreateEventRequest

	userDataCtx, err := utils.GetUserDataFromCtx(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "parsing_ctx_error",
			Message: err.Error(),
		})
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	newEvent, err := h.service.CreateEvent(&req, userDataCtx)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "create_event_error",
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusCreated, newEvent)
}
