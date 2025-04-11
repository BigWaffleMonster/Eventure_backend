package category

import "github.com/jinzhu/copier"

type CategoryService interface {
	GetCategoryList() (*[]CategoryView, error)
	GetCategoryByID(id uint) (*CategoryView, error)
}

type categoryService struct {
	Repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryService{Repo: repo}
}

func (s *categoryService) GetCategoryList() (*[]CategoryView, error) {
	var categoryView []CategoryView

	data, err := s.Repo.GetCategoryList()
	if err != nil {
		return nil, err
	}

	copier.Copy(&categoryView, &data)

	return &categoryView, nil
}

func (s *categoryService) GetCategoryByID(id uint) (*CategoryView, error) {
	var categoryView CategoryView

	data, err := s.Repo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	copier.Copy(&categoryView, &data)

	return &categoryView, nil
}
