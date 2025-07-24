package interfaces

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"gorm.io/gorm"
)

type IUnitOfWork[TRepo IBaseRepository[TEntity], TEntity any] interface{
	RunInTx(
		repository func(tx *gorm.DB) TRepo,
		execute func(repo TRepo, store DomainEventStore) results.Result) results.Result
	Repository() TRepo
	DomainEventStore() DomainEventStore
}