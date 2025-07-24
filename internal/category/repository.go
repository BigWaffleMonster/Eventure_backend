package category

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	interfaces.IBaseRepository[Category]
}

type categoryRepository struct {
	repository.BaseRepository[Category]
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		repository.BaseRepository[Category]{DB: db},
	}
}