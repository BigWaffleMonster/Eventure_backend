package domain_events_handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/google/uuid"
)

type participantCreatedDomainEventHandler struct{
	EventRepo event.EventRepository
	ParticipantRepo participant.ParticipantRepository
}

func NewParticipantCreatedDomainEventHandler(eRepo event.EventRepository, pRepo participant.ParticipantRepository) domain_events_abstractions.DomainEventHandler{
	return &participantCreatedDomainEventHandler{
		EventRepo: eRepo,
		ParticipantRepo: pRepo,
	}
}

func (h * participantCreatedDomainEventHandler) Handle(domainEventData *domain_events_abstractions.DomainEventData) error {

	if domainEventData.Type != "ParticipantCreatedDomainEvent" {
		return fmt.Errorf("failed to handler event. event type is incorrect")
	}

	var domainEvent domain_events.ParticipantCreatedDomainEvent
	err := json.Unmarshal([]byte(domainEventData.Content) , &domainEvent)
	if err != nil {
		log.Print(err)
	}

	var event *event.Event

	event, err = h.EventRepo.GetByID(domainEvent.EventID)
	if err != nil {
		return err
	}

	var participants *[]participant.Participant

	participants, err = h.ParticipantRepo.GetCollection(domainEvent.EventID)

	if err != nil {
		return err
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
		return nil
	}

	err = h.EventRepo.Update(event)

	if err != nil {
		return err
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
	return domainEventData.Type == "ParticipantCreatedDomainEvent"
}