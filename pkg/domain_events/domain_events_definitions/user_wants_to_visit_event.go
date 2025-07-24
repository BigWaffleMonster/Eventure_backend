package domain_events_definitions

import (
	"encoding/json"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/BigWaffleMonster/Eventure_backend/utils/enums"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type UserWantsToVisitEvent struct{
	EventID uuid.UUID
	UserID uuid.UUID
	ActualQTYOfGuests int
	Status string
}

func NewUserWantsToVisitEvent(
	eventId uuid.UUID, 
	userID uuid.UUID, 
	actualQTYOfGuests int,
	status string) (*domain_events_base.DomainEventData, results.Result){
		
	event := UserWantsToVisitEvent{
		EventID: eventId,
		UserID: userID,
		Status: status,
		ActualQTYOfGuests: actualQTYOfGuests,
	}

    b, err := json.Marshal(event)
    if err != nil {
        return nil, results.NewInternalError(err.Error())
    }

	domainEventData := domain_events_base.NewDomainEventData(enums.UserWantsToVisitEvent, string(b))

	return &domainEventData, results.NewResultOk()
}