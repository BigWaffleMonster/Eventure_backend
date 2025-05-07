package category

import "github.com/jinzhu/copier"

type CategoryService interface {
	GetCollection() (*[]CategoryView, error)
	GetByID(id uint) (*CategoryView, error)
}

type categoryService struct {
	Repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryService{Repo: repo}
}

func (s *categoryService) GetCollection() (*[]CategoryView, error) {
	var categoryView []CategoryView

	data, err := s.Repo.GetCollection()
	if err != nil {
		return nil, err
	}

	copier.Copy(&categoryView, &data)

	return &categoryView, nil
}

func (s *categoryService) GetByID(id uint) (*CategoryView, error) {
	var categoryView CategoryView

	data, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	copier.Copy(&categoryView, &data)

	return &categoryView, nil
}
