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

func (h *ParticipantHandler) GetParticipantsFromEvent(c *gin.Context) {
	eventID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	participants, err := h.service.GetParticipantsFromEvent(eventID)
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	global_utils.SendSuccess(c, participants, "")
}

func (r *ParticipantHandler) AddParticipantToEvent(userID, eventID uuid.UUID) error {
	return nil
}

func (r *ParticipantHandler) RemoveParticipantFromEvent(userID, eventID uuid.UUID) error {
	return nil
}

func (r *ParticipantHandler) RemoveAllParticipantsFromEvent(eventID uuid.UUID) error {
	return nil
}
