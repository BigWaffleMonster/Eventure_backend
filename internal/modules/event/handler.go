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

func (h *EventHandler) GetEvents(c *gin.Context) {
	events, err := h.service.GetEvents()
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	global_utils.SendSuccess(c, events, "")
}

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

func (h *EventHandler) GetUserCreatedEvents(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	event, err := h.service.GetUserCreatedEvents(userID)
	if err != nil {
		global_utils.SendError(c, err)
		return
	}

	global_utils.SendSuccess(c, event, "")
}

func (h *EventHandler) RemoveEvent(c *gin.Context) {
	eventID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	err = h.service.RemoveEvent(eventID)
	if err != nil {
		global_utils.SendError(c, err)
		return
	}

	global_utils.SendSuccess(c, "", "Event removed")
}

func (h *EventHandler) UpdateEvent(c *gin.Context) {
	eventID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	userDataCtx, err := utils.GetUserDataFromCtx(c)
	if err != nil {
		global_utils.SendError(c, global_utils.NewAppErrorWithErr(
			http.StatusBadRequest,
			"Ошибка контекста",
			err,
		))

		return
	}

	var req UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global_utils.SendError(c, global_utils.NewAppErrorWithErr(
			http.StatusBadRequest,
			"Ошибка парсинга тела запроса",
			err,
		))
		return
	}

	err = h.service.UpdateEvent(eventID, userDataCtx, &req)
	if err != nil {
		global_utils.SendError(c, err)
		return
	}

	global_utils.SendSuccess(c, "", "Event updated")
}
