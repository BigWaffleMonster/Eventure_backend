package participant

import (
	global_utils "github.com/BigWaffleMonster/Eventure_backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// utils "github.com/BigWaffleMonster/Eventure_backend/internal/modules/event/utils"
// "github.com/gin-gonic/gin"
// "github.com/google/uuid"

type ParticipantHandler struct {
	service *ParticipantService
}

func NewParticipantHandler(service *ParticipantService) *ParticipantHandler {
	return &ParticipantHandler{service: service}
}

// TODO test with postman
func (h *ParticipantHandler) GetParticipantsFromEvent(c *gin.Context) {
	eventID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	participants, err := h.service.GetParticipantsFromEvent(eventID)
	if err != nil {
		global_utils.SendError(c, err)
		return
	}

	global_utils.SendSuccess(c, participants, "")
}

// TODO test with postman and go through code one more time
func (h *ParticipantHandler) AddParticipantToEvent(c *gin.Context) {
	eventID, err := uuid.Parse(c.Query("event_id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	userID, err := uuid.Parse(c.Query("user_id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	err = h.service.AddParticipantToEvent(userID, eventID)
	if err != nil {
		global_utils.SendError(c, err)
		return
	}

	global_utils.SendSuccess(c, "", "Successfully added to event!")
}

func (h *ParticipantHandler) RemoveParticipantFromEvent(c *gin.Context) {
	eventID, err := uuid.Parse(c.Query("event_id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	userID, err := uuid.Parse(c.Query("user_id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	err = h.service.RemoveParticipantFromEvent(userID, eventID)
	if err != nil {
		global_utils.SendError(c, err)
		return
	}

	global_utils.SendSuccess(c, "", "Successfully removed from event!")
}

func (h *ParticipantHandler) RemoveAllParticipantsFromEvent(c *gin.Context) {
	eventID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	err = h.service.RemoveAllParticipantsFromEvent(eventID)
	if err != nil {
		global_utils.SendError(c, err)
		return
	}

	global_utils.SendSuccess(c, "", "Successfully removed all participants from event!")
}
