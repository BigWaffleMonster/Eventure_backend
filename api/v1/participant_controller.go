package v1

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ParticipantController interface {
	Create(ctx *gin.Context)
	ChangeState(ctx *gin.Context)
	Remove(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetCollection(ctx *gin.Context)
}

type participantController struct {
	Service participant.ParticipantService
}

func NewParticipantController(service participant.ParticipantService) ParticipantController {
	return &participantController{Service: service}
}

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
func (c *participantController) Create(ctx *gin.Context) {
	var body participant.ParticipantInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.Service.Create(&body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

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
func (c *participantController) ChangeState(ctx *gin.Context) {
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

	err = c.Service.ChangeState(id, &body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

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
func (c *participantController) Remove(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.Service.Remove(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

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
func (c *participantController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	participantView, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, participantView)
}

// @summary Получение участников события
// @schemes
// @description Получение участников события
// @tags participant
// @accept json
// @produce json
// @param id path string true "Идентификатор события"
// @success 200 {object} []participant.ParticipantView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /participant [get]
func (c *participantController) GetCollection(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	participantViews, err := c.Service.GetCollection(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, participantViews)
}