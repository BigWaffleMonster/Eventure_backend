package repository

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository[T any] interface{
	Create(event *T ) results.Result
	Delete(id uuid.UUID) results.Result
	Update(event *T ) results.Result
	GetByID(id interface{}) (*T , results.Result)
	GetCollection() (*[]T , results.Result)
	GetByExpression(expr string, conds ...any) (*T , results.Result)
	GetCollectionByExpression(expr string, conds ...any) (*[]T, results.Result)
}

type RepositoryImp[T any] struct{
	DB *gorm.DB
}

func NewRepository[T any] (db *gorm.DB) Repository[T]  {
	return &RepositoryImp[T] {DB: db}
}

func (r *RepositoryImp[T]) Create(entity *T) results.Result {
	err := r.DB.Create(entity).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *RepositoryImp[T]) Delete(id uuid.UUID) results.Result {
	var entity T

	err := r.DB.Where("id = ?", id).Delete(&entity).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *RepositoryImp[T]) Update(entity *T) results.Result {
	err := r.DB.Save(entity).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *RepositoryImp[T]) GetByID(id interface{}) (*T, results.Result) {
	var entity T

	result := r.DB.First(&entity, "id = ?", id)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &entity, results.NewResultOk()
}

func (r *RepositoryImp[T]) GetCollection() (*[]T, results.Result){
	var entities []T

	result := r.DB.Find(&entities)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &entities, results.NewResultOk()
}

func (r *RepositoryImp[T]) GetByExpression(expr string, conds ...any) (*T, results.Result){
	var entity T

	result := r.DB.First(&entity, expr, conds)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &entity, results.NewResultOk()
}


func (r *RepositoryImp[T]) GetCollectionByExpression(expr string, conds ...any) (*[]T, results.Result){
	var entities []T

	result := r.DB.Find(&entities, expr, conds)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &entities, results.NewResultOk()
}