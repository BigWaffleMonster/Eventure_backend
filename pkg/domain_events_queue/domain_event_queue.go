package domain_events_queue

import (
	"fmt"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
)

type domainEventQueue struct{
	domainEventBus domain_events_abstractions.DomainEventBus
	CancelationRequested bool
}

func NewDomainEventQueue(domainEventBus domain_events_abstractions.DomainEventBus) domain_events_abstractions.DomainEventQueue{
	return &domainEventQueue{
		domainEventBus: domainEventBus,
		CancelationRequested: false,
	}
}

func (p *domainEventQueue) StartQueue() {
	go p.startQueue()
}

func (p *domainEventQueue) startQueue() {
	p.CancelationRequested = false

	for {
		p.domainEventBus.PublishAll()

		if (p.CancelationRequested){
			break
		}

		fmt.Println("Queue working, time: ", time.Now().UTC())
		time.Sleep(2 * time.Minute)
	}
}

func (p *domainEventQueue) StopQueue() {
	p.CancelationRequested = true
}