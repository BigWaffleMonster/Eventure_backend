package unit_of_work

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"gorm.io/gorm"
)

type UnitOfWork[TRepo interfaces.IBaseRepository[TEntity], TEntity any] struct{
	DB *gorm.DB
	Repo TRepo
	Store interfaces.DomainEventStore
}

func (uof *UnitOfWork[TRepo, TEntity]) RunInTx(
	ctx context.Context, 
	repository func(tx *gorm.DB) TRepo,
	execute func(repo TRepo, store interfaces.DomainEventStore) results.Result) results.Result {
	tx := uof.DB.Begin()

	repo := repository(tx)
	store := domain_events.NewDomainEventStore(tx)

	result := execute(repo, store)

	if result.IsSuccess {
		tx.Commit()
		return result
	}

	tx.Rollback()

	return result
}

func (uof *UnitOfWork[TRepo, TEntity]) Repository(ctx context.Context) TRepo{
	return uof.Repo
}

func (uof *UnitOfWork[TRepo, TEntity]) DomainEventStore(ctx context.Context) interfaces.DomainEventStore{
	return uof.Store
}