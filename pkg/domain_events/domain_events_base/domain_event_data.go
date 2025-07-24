package domain_events_base

import (
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/utils/enums"
	"github.com/google/uuid"
)

type DomainEventData struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid"`
	Type enums.EventType 
	Content string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func NewDomainEventData(eventType enums.EventType, content string) DomainEventData{
	return DomainEventData{
		ID: uuid.New(),
		Type: eventType,
		Content: content,
	}
}