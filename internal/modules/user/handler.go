package user

import (
	global_utils "github.com/BigWaffleMonster/Eventure_backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		global_utils.SendError(c, global_utils.ErrBadRequest)
		return
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		global_utils.SendError(c, err)
		return
	}

	global_utils.SendSuccess(c, user, "")
}

func (h *UserHandler) GetUserByEmail() {}
func (h *UserHandler) UpdateUser()     {}
func (h *UserHandler) RemoveUser()     {}
