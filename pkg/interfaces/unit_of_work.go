package interfaces

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"gorm.io/gorm"
)

type IUnitOfWork[TRepo IBaseRepository[TEntity], TEntity any] interface{
	RunInTx(
		ctx context.Context, 
		repository func(tx *gorm.DB) TRepo,
		execute func(repo TRepo, store DomainEventStore) results.Result) results.Result
	Repository(ctx context.Context) TRepo
	DomainEventStore(ctx context.Context) DomainEventStore
}