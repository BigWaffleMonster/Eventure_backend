package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
type ParticipantController struct {
	Service participant.ParticipantService
}

func NewParticipantController(service participant.ParticipantService) *ParticipantController {
	return &ParticipantController{Service: service}
}

// @Security ApiKeyAuth
// @summary Создание участника
// @schemes
// @description Создание участника
// @tags participant
// @accept json
// @produce json
// @param participant body participant.ParticipantInput false "Участник"
// @success 204 
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /participant [post]
func (c *ParticipantController) Create(ctx *gin.Context) {
	var body participant.ParticipantInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := ctx.MustGet(auth.CurrentUserVarName).(*auth.CurrentUser)

	err := c.Service.Create(&body, currentUser)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @Security ApiKeyAuth
// @summary Обновить участника
// @schemes
// @description Обновить участника
// @tags participant
// @accept json
// @produce json
// @param id path string true "Идентификатор участника"
// @param participant body participant.ParticipantInput false "Участник"
// @success 204 
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /participant/{id} [put]
func (c *ParticipantController) ChangeState(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body string

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := ctx.MustGet(auth.CurrentUserVarName).(auth.CurrentUser)

	err = c.Service.ChangeState(id, &body, currentUser)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @Security ApiKeyAuth
// @summary Удаление участника
// @schemes
// @description Удаление участника
// @tags participant
// @accept json
// @produce json
// @param id path string true "Идентификатор участника"
// @success 204 
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /participant/{id} [delete]
func (c *ParticipantController) Remove(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := ctx.MustGet(auth.CurrentUserVarName).(auth.CurrentUser)

	err = c.Service.Remove(id, currentUser)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @Security ApiKeyAuth
// @summary Получение участника
// @schemes
// @description Получение участника
// @tags participant
// @accept json
// @produce json
// @param id path string true "Идентификатор участника"
// @success 200 {object} participant.ParticipantView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /participant/{id} [get]
func (c *ParticipantController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := ctx.MustGet(auth.CurrentUserVarName).(auth.CurrentUser)

	participantView, err := c.Service.GetByID(id, currentUser)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, participantView)
}

// @Security ApiKeyAuth
// @summary Получение участников события
// @schemes
// @description Получение участников события
// @tags participant
// @accept json
// @produce json
// @param eventId path string true "Идентификатор события"
// @success 200 {object} []participant.ParticipantView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /participant/event [get]
func (c *ParticipantController) GetCollection(ctx *gin.Context) {
	eventID, err := uuid.Parse(ctx.Param("eventI"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	participantViews, err := c.Service.GetCollection(eventID)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, participantViews)
}

// @Security ApiKeyAuth
// @summary Получение своих записей
// @schemes
// @description Получение своих записей
// @tags participant
// @accept json
// @produce json
// @success 200 {object} []participant.ParticipantView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /participant [get]
func (c *ParticipantController) GetOwnedCollection(ctx *gin.Context) {
	currentUser := ctx.MustGet(auth.CurrentUserVarName).(*auth.CurrentUser)

	participantViews, err := c.Service.GetOwnedCollection(currentUser)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, participantViews)
}