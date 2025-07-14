package v1

import (
	"net/http"
	"strconv"

	"github.com/BigWaffleMonster/Eventure_backend/internal/category"
	"github.com/BigWaffleMonster/Eventure_backend/utils/responses"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	Service category.CategoryService
}

// @summary Получение категорий
// @schemes
// @description Получение категорий
// @tags category
// @accept json
// @produce json
// @success 200 {object} responses.ResponseOk[[]category.CategoryView]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /category [get]
func NewCategoryController(service category.CategoryService) *CategoryController {
	return &CategoryController{Service: service}
}

func (c *CategoryController) GetCollection(ctx *gin.Context) {
	resp, result := c.Service.GetCollection()
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to get collection", result.Errors))
		return
	}

	ctx.JSON(http.StatusOK, responses.NewResponseOk(&resp, "Get collection success"))
}

// @summary Получение категории
// @schemes
// @description Получение категории
// @tags category
// @accept json
// @produce json
// @param id path string true "Идентификатор категории"
// @success 200 {object} responses.ResponseOk[category.CategoryView]
// @failure 400 {object} responses.ResponseFailed
// @failure 401 {object} responses.ResponseFailed
// @failure 403 {object} responses.ResponseFailed
// @failure 404 {object} responses.ResponseFailed
// @failure 409 {object} responses.ResponseFailed
// @failure 500 {object} responses.ResponseFailed
// @router /category/{id} [get]
func (c *CategoryController) GetByID(ctx *gin.Context) {
	str_id := ctx.Query("id")

	u64, err := strconv.ParseUint(str_id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewResponseFailed("Failed to get user", []string{err.Error()}))
		return
	}

	id := uint(u64)

	resp, result := c.Service.GetByID(id)
	if result.IsFailed {
		ctx.JSON(result.Code, responses.NewResponseFailed("Failed to get user", result.Errors))
		return
	}

	ctx.JSON(http.StatusOK, responses.NewResponseOk(&resp, "Get user success"))
}
