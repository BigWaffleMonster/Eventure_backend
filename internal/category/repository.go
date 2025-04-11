package category

import (
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoryList() (*[]Category, error)
	GetCategoryByID(id uint) (*Category, error)
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{DB: db}
}

func (r *categoryRepository) GetCategoryList() (*[]Category, error) {
	var category []Category
	result := r.DB.Find(&category)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (r *categoryRepository) GetCategoryByID(id uint) (*Category, error) {
	var category Category
	result := r.DB.Where("id = ?", id).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
