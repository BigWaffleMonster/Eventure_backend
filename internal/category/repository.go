package category

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCollection() (*[]Category, results.Result)
	GetByID(id uint) (*Category, results.Result)
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{DB: db}
}

func (r *categoryRepository) GetCollection() (*[]Category, results.Result) {
	var category []Category
	result := r.DB.Find(&category)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &category, results.NewResultOk()
}

func (r *categoryRepository) GetByID(id uint) (*Category, results.Result) {
	var category Category

	result := r.DB.Where("id = ?", id).First(&category)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &category, results.NewResultOk()
}
