package domain_events_definitions

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/utils/enums"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type UserCanVisitEvent struct{
	EventID uuid.UUID
	UserID uuid.UUID
	Status string
}

func NewUserCanVisitEvent(
	eventId uuid.UUID, 
	userID uuid.UUID, 
	status string) (*domain_events_base.DomainEventData, results.Result){
		
	event := UserCanVisitEvent{
		EventID: eventId,
		UserID: userID,
		Status: status,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, results.NewInternalError(err.Error())
    }

	domainEventData := domain_events_base.NewDomainEventData(enums.UserCanVisitEvent, string(b))

	return &domainEventData, results.NewResultOk()
}