package category

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"gorm.io/gorm"
)

func NewCategoryRepository(db *gorm.DB) repository.Repository[Category] {
	return repository.NewRepository[Category](db)
}
