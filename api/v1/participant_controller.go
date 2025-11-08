package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/utils/responses"
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
// @success 201 {object} responses.ResponseOkString
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /participant [post]
func (c *ParticipantController) Create(ctx *gin.Context) {
	var body participant.ParticipantInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to create participant", []string{err.Error()}))
		return
	}

	result := c.Service.Create(ctx.Request.Context(), &body)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to create participant", result.Errors))
		return
	}

	ctx.JSON(http.StatusCreated, responses.NewResponseOkString("Created participant success"))
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
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /participant/{id} [put]
func (c *ParticipantController) ChangeState(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to update participant", []string{err.Error()}))
		return
	}

	var body string

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to update participant", []string{err.Error()}))
		return
	}

	result := c.Service.ChangeState(ctx.Request.Context(), id, &body)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to update participant", result.Errors))
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
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /participant/{id} [delete]
func (c *ParticipantController) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to delete participant", []string{err.Error()}))
		return
	}

	result := c.Service.Delete(ctx.Request.Context(), id)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to delete participant", result.Errors))
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
// @success 200 {object} responses.ResponseOk[participant.ParticipantView]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /participant/{id} [get]
func (c *ParticipantController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to get participant", []string{err.Error()}))
		return
	}

	participantView, result := c.Service.GetByID(ctx.Request.Context(), id)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to get participant", result.Errors))
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, responses.NewResponseOk(&participantView, "Get participant success"))
}

// @Security ApiKeyAuth
// @summary Получение участников события
// @schemes
// @description Получение участников события
// @tags participant
// @accept json
// @produce json
// @param eventId path string true "Идентификатор события"
// @success 200 {object} responses.ResponseOk[[]participant.ParticipantView]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /participant/event/{eventId} [get]
func (c *ParticipantController) GetCollection(ctx *gin.Context) {
	eventID, err := uuid.Parse(ctx.Param("eventId"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to get participants", []string{err.Error()}))
		return
	}

	participantViews, result := c.Service.GetCollection(ctx.Request.Context(), eventID)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to get participants", result.Errors))
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, responses.NewResponseOk(&participantViews, "Get participants success"))
}

// @Security ApiKeyAuth
// @summary Получение своих записей
// @schemes
// @description Получение своих записей
// @tags participant
// @accept json
// @produce json
// @success 200 {object} responses.ResponseOk[[]participant.ParticipantView]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /participant [get]
func (c *ParticipantController) GetOwnedCollection(ctx *gin.Context) {
	participantViews, result := c.Service.GetOwnedCollection(ctx.Request.Context())
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to get participants", result.Errors))
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, responses.NewResponseOk(&participantViews, "Get participants success"))
}