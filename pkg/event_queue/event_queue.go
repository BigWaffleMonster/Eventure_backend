package publisher

import (
	"fmt"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
)

type eventQueue struct{
	EventBus domain_events_abstractions.EventBus
	CancelationRequested bool
}

func NewPublisher(eventBus domain_events_abstractions.EventBus) domain_events_abstractions.EventQueue{
	return &eventQueue{
		EventBus: eventBus,
		CancelationRequested: false,
	}
}

func (p *eventQueue) StartQueue() {
	go p.startQueue()
}

func (p *eventQueue) startQueue() {
	 for {
		p.handle()

		if (p.CancelationRequested){
			break
		}

		time.Sleep(2 * time.Minute)
	}
}

func (p *eventQueue) StopQueue() {
	p.CancelationRequested = true
}

func (p *eventQueue) handle() {
	domainEvents, err := p.EventBus.GetDomainEvents()

	if err != nil {
		fmt.Printf("Failed to get domain events, %e\n", err)
	}

	for _, domainEvent := range domainEvents{
		err = p.EventBus.Publish(domainEvent)

		if err != nil {
			fmt.Printf("Failed to publish domain event: %s, %e\n", domainEvent.ID.String(), err)
		}

		err = p.EventBus.RemoveFromStore(domainEvent)

		if err != nil {
			fmt.Printf("Failed to remove domain event: %s, %e\n", domainEvent.ID.String(), err)
		}
	}
}