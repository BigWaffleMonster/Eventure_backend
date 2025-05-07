package category

import (
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCollection() (*[]Category, error)
	GetByID(id uint) (*Category, error)
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{DB: db}
}

func (r *categoryRepository) GetCollection() (*[]Category, error) {
	var category []Category
	result := r.DB.Find(&category)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (r *categoryRepository) GetByID(id uint) (*Category, error) {
	var category Category
	result := r.DB.Where("id = ?", id).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
