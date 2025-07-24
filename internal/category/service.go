package category

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CategoryService interface {
	GetCollection() (*[]CategoryView, results.Result)
	GetByID(id uuid.UUID) (*CategoryView, results.Result)
}

type categoryService struct {
	Repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryService{Repo: repo}
}

func (s *categoryService) GetCollection() (*[]CategoryView, results.Result) {
	var categoryView []CategoryView

	data, result := s.Repo.GetCollection()

	if result.IsFailed {
		return nil, result
	}

	copier.Copy(&categoryView, &data)

	return &categoryView, results.NewResultOk()
}

func (s *categoryService) GetByID(id uuid.UUID) (*CategoryView, results.Result) {
	var categoryView CategoryView

	data, result := s.Repo.GetByID(id)

	if result.IsFailed {
		return nil, result
	}

	if data == nil {
		return nil, results.NewNotFoundError("Category")
	}

	copier.Copy(&categoryView, &data)

	return &categoryView, results.NewResultOk()
}
