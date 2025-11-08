package domain_events

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"gorm.io/gorm"
)

type domainEventStore struct{
	DB *gorm.DB
}

func NewDomainEventStore(db *gorm.DB) interfaces.DomainEventStore{
	return &domainEventStore{
		DB: db,
	}
}

func (e *domainEventStore) AddToStore(ctx context.Context, domainEventData *domain_events_base.DomainEventData) results.Result{
	result := e.DB.Create(domainEventData).Error

	if result != nil {
		return results.NewInternalError(result.Error())
	}

	return results.NewResultOk()
}