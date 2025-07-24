package domain_events_definitions

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/utils/enums"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type UserDeleted struct{
	UserID uuid.UUID
}

func NewUserDeleted(userID uuid.UUID) (*domain_events_base.DomainEventData, results.Result){
	event := UserDeleted{
		UserID: userID,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, results.NewInternalError(err.Error())
    }

	domainEventData := domain_events_base.NewDomainEventData(enums.UserDeleted, string(b))

	return &domainEventData, results.NewResultOk()
}