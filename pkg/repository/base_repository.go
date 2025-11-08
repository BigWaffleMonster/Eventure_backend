package repository

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct{
	DB *gorm.DB
}

func NewRepository[T any] (db *gorm.DB) interfaces.IBaseRepository[T]  {
	return &BaseRepository[T] {DB: db}
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) results.Result {
	err := r.DB.Create(entity).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id uuid.UUID) results.Result {
	var entity T

	err := r.DB.Where("id = ?", id).Delete(&entity).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) results.Result {
	err := r.DB.Save(entity).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *BaseRepository[T]) GetByID(ctx context.Context, id interface{}) (*T, results.Result) {
	var entity T

	result := r.DB.First(&entity, "id = ?", id)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &entity, results.NewResultOk()
}

func (r *BaseRepository[T]) GetCollection(ctx context.Context) (*[]T, results.Result){
	var entities []T

	result := r.DB.Find(&entities)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &entities, results.NewResultOk()
}

func (r *BaseRepository[T]) GetByExpression(ctx context.Context, expr string, conds ...any) (*T, results.Result){
	var entity T

	result := r.DB.First(&entity, expr, conds)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &entity, results.NewResultOk()
}


func (r *BaseRepository[T]) GetCollectionByExpression(ctx context.Context, expr string, conds ...any) (*[]T, results.Result){
	var entities []T

	result := r.DB.Find(&entities, expr, conds)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &entities, results.NewResultOk()
}