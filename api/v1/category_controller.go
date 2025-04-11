package v1

import (
	"net/http"
	"strconv"

	"github.com/BigWaffleMonster/Eventure_backend/internal/category"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetCategoryList(ctx *gin.Context)
	GetCategoryByID(ctx *gin.Context)
}

type categoryController struct {
	Service category.CategoryService
}

func NewCategoryController(service category.CategoryService) CategoryController {
	return &categoryController{Service: service}
}

func (c *categoryController) GetCategoryList(ctx *gin.Context) {
	resp, err := c.Service.GetCategoryList()
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": resp})
}

func (c *categoryController) GetCategoryByID(ctx *gin.Context) {
	str_id := ctx.Query("id")

	u64, err := strconv.ParseUint(str_id, 10, 32)
	id := uint(u64)

	resp, err := c.Service.GetCategoryByID(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": resp})
}
