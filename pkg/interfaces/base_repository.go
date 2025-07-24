package interfaces

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type IBaseRepository[T any] interface{
	Create(event *T ) results.Result
	Delete(id uuid.UUID) results.Result
	Update(event *T ) results.Result
	GetByID(id interface{}) (*T , results.Result)
	GetCollection() (*[]T , results.Result)
	GetByExpression(expr string, conds ...any) (*T , results.Result)
	GetCollectionByExpression(expr string, conds ...any) (*[]T, results.Result)
}