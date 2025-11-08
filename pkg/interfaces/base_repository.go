package interfaces

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type IBaseRepository[T any] interface{
	Create(ctx context.Context, event *T ) results.Result
	Delete(ctx context.Context, id uuid.UUID) results.Result
	Update(ctx context.Context, event *T ) results.Result
	GetByID(ctx context.Context, id interface{}) (*T , results.Result)
	GetCollection(ctx context.Context) (*[]T , results.Result)
	GetByExpression(ctx context.Context, expr string, conds ...any) (*T , results.Result)
	GetCollectionByExpression(ctx context.Context, expr string, conds ...any) (*[]T, results.Result)
}