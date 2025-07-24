package domain_events_definitions

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/utils/enums"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type EventDeleted struct{
	EventID uuid.UUID
}

func NewEventDeleted(eventId uuid.UUID) (*domain_events_base.DomainEventData, results.Result){
	event := EventDeleted{
		EventID: eventId,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, results.NewInternalError(err.Error())
    }

	domainEventData := domain_events_base.NewDomainEventData(enums.EventDeleted, string(b))

	return &domainEventData, results.NewResultOk()
}
