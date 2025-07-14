package domain_events

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/google/uuid"
)

func NewDomainEventData(eventType string, content string) domain_events_abstractions.DomainEventData{
	return domain_events_abstractions.DomainEventData{
		ID: uuid.New(),
		Type: eventType,
		Content: content,
	}
}