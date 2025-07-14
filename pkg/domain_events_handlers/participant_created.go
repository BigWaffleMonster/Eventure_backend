package domain_events_handlers

import (
	"encoding/json"
	"fmt"

	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

const participantCreatedDomainType = "ParticipantCreatedDomainEvent"

type participantCreatedDomainEventHandler struct{
	EventRepo event.EventRepository
	ParticipantRepo participant.ParticipantRepository
}

func NewParticipantCreatedDomainEventHandler(
	eRepo event.EventRepository, 
	pRepo participant.ParticipantRepository) domain_events_abstractions.DomainEventHandler{

	return &participantCreatedDomainEventHandler{
		EventRepo: eRepo,
		ParticipantRepo: pRepo,
	}
}

func (h * participantCreatedDomainEventHandler) Handle(domainEventData *domain_events_abstractions.DomainEventData) results.Result {

	if domainEventData.Type != participantCreatedDomainType {
		return results.NewInvalidDomainTypeError(domainEventData.Type, participantCreatedDomainType)
	}

	var domainEvent domain_events.ParticipantCreatedDomainEvent
	err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
    if err != nil {
        return results.NewInternalError(err.Error())
    }

	var event *event.Event

	event, result := h.EventRepo.GetByID(domainEvent.EventID)
	if result.IsFailed {
		return result
	}

	var participants *[]participant.Participant

	participants, result = h.ParticipantRepo.GetCollection(domainEvent.EventID)

	if result.IsFailed {
		return result
	}

	var participantsWhoGo []participant.Participant

	for _, p := range *participants{
		if p.Status != "No" {
			participantsWhoGo = append(participantsWhoGo, p)
		}
	}

	if event.MaxQtyParticipants == len(participantsWhoGo){
		//TODO: Notify user
		fmt.Println("failed to create new participant, max QTY of participants reached")
		return results.NewResultOk()
	}

	result = h.EventRepo.Update(event)

	if result.IsFailed {
		return result
	}

	participant := participant.Participant{
		ID: uuid.New(),
		UserID: domainEvent.UserID,
		EventID: domainEvent.EventID,
		Status: domainEvent.Status,
		Ticket: domainEvent.Ticket,
	}
	
	return h.ParticipantRepo.Create(&participant)
}

func (h * participantCreatedDomainEventHandler) IsTypeOf(domainEventData *domain_events_abstractions.DomainEventData) bool {
	return domainEventData.Type == participantCreatedDomainType
}