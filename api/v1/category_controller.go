package v1

import (
	"net/http"
	"strconv"

	"github.com/BigWaffleMonster/Eventure_backend/internal/category"
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
// @success 200 {object} []category.CategoryView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /category [get]
func NewCategoryController(service category.CategoryService) *CategoryController {
	return &CategoryController{Service: service}
}

func (c *CategoryController) GetCollection(ctx *gin.Context) {
	resp, err := c.Service.GetCollection()
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @summary Получение категории
// @schemes
// @description Получение категории
// @tags category
// @accept json
// @produce json
// @param id path string true "Идентификатор категории"
// @success 200 {object} category.CategoryView
// @failure 400 {string} string "error"
// @failure 409 {string} string "error"
// @failure 500 {string} string "error"
// @router /category/{id} [get]
func (c *CategoryController) GetByID(ctx *gin.Context) {
	str_id := ctx.Query("id")

	u64, err := strconv.ParseUint(str_id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(u64)

	resp, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
