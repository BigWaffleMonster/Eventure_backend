package domain_events

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type UserDeletedDomainEvent struct{
	UserID uuid.UUID
}

func NewUserDeletedDomainEvent(userID uuid.UUID) (*domain_events_abstractions.DomainEventData, results.Result){
	event := UserDeletedDomainEvent{
		UserID: userID,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, results.NewInternalError(err.Error())
    }

	domainEventData := NewDomainEventData("UserDeletedDomainEvent", string(b))

	return &domainEventData, results.NewResultOk()
}