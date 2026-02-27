package event

import (
	"net/http"

	utils "github.com/BigWaffleMonster/Eventure_backend/internal/modules/event/utils"
	global_utils "github.com/BigWaffleMonster/Eventure_backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		global_utils.SendError(c, global_utils.NewAppErrorWithErr(
			http.StatusBadRequest,
			"Ошибка контекста",
			err,
		))

		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		global_utils.SendError(c, global_utils.NewAppErrorWithErr(
			http.StatusBadRequest,
			"Ошибка валидации",
			err,
		))

		return
	}

	newEvent, err := h.service.CreateEvent(&req, userDataCtx)
	if err != nil {
		global_utils.SendError(c, err)

		return
	}

	global_utils.SendSuccessWithStatus(c, http.StatusCreated, newEvent, "Событие создано")
}

func (h *EventHandler) GetEvents(c *gin.Context) {}

func (h *EventHandler) GetEventByID(c *gin.Context) {
	eventID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	event, err := h.service.GetEventByID(eventID)
	if err != nil {
		global_utils.SendError(c, err)
		return
	}

	global_utils.SendSuccess(c, event, "")
}
