package user

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/unit_of_work"
	"gorm.io/gorm"
)

type UnitOfWork interface{
	interfaces.IUnitOfWork[UserRepository, User]
}

type unitOfWork struct{
	unit_of_work.UnitOfWork[UserRepository, User]
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork{
	return &unitOfWork{
		unit_of_work.UnitOfWork[UserRepository, User]{
			DB: db,
			Repo: NewUserRepository(db),
			Store: domain_events.NewDomainEventStore(db),
		},
	}
}