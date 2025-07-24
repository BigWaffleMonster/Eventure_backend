package event

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/unit_of_work"
	"gorm.io/gorm"
)

type UnitOfWork interface{
	interfaces.IUnitOfWork[EventRepository, Event]
}

type unitOfWork struct{
	unit_of_work.UnitOfWork[EventRepository, Event]
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork{
	return &unitOfWork{
		unit_of_work.UnitOfWork[EventRepository, Event]{
			DB: db,
			Repo: NewEventRepository(db),
			Store: domain_events.NewDomainEventStore(db),
		},
	}
}