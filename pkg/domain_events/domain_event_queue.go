package domain_events

import (
	"fmt"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
)

type domainEventQueue struct{
	domainEventBus interfaces.DomainEventBus
	CancelationRequested bool
}

func NewDomainEventQueue(domainEventBus interfaces.DomainEventBus) interfaces.DomainEventQueue{
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
		p.domainEventBus.Publish()

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