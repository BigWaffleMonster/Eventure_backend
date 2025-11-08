package category

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/fx"
)

type CategoryService interface {
	GetCollection(ctx context.Context) (*[]CategoryView, results.Result)
	GetByID(ctx context.Context, id uuid.UUID) (*CategoryView, results.Result)
}

type CategoryServiceParams struct {
	fx.In

	Repo CategoryRepository
}

type categoryService struct {
	repo CategoryRepository
}

func NewCategoryService(p CategoryServiceParams) CategoryService {
	return &categoryService{repo: p.Repo}
}

func (s *categoryService) GetCollection(ctx context.Context) (*[]CategoryView, results.Result) {
	var categoryView []CategoryView

	data, result := s.repo.GetCollection(ctx)

	if result.IsFailed {
		return nil, result
	}

	copier.Copy(&categoryView, &data)

	return &categoryView, results.NewResultOk()
}

func (s *categoryService) GetByID(ctx context.Context, id uuid.UUID) (*CategoryView, results.Result) {
	var categoryView CategoryView

	data, result := s.repo.GetByID(ctx, id)

	if result.IsFailed {
		return nil, result
	}

	if data == nil {
		return nil, results.NewNotFoundError("Category")
	}

	copier.Copy(&categoryView, &data)

	return &categoryView, results.NewResultOk()
}
