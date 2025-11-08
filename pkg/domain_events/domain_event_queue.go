package domain_events

import (
	"context"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	"go.uber.org/fx"
)

type domainEventQueue struct{
	domainEventBus interfaces.DomainEventBus
	logger         sglogger.Logger

	ctx            context.Context
	cancel         context.CancelFunc
	isRunning      bool
	startReady     chan struct{}
}

type DomainEventQueueParams struct {
	fx.In

	DomainEventBus interfaces.DomainEventBus
	Logger         sglogger.Logger
}

func NewDomainEventQueue(p DomainEventQueueParams) interfaces.DomainEventQueue{
	ctx, cancel := context.WithCancel(context.Background())

	return &domainEventQueue{
		domainEventBus: p.DomainEventBus,
		logger:    p.Logger,
		ctx:       ctx,
		cancel:    cancel,
		isRunning: false,
		startReady: make(chan struct{}),
	}
}

func (p *domainEventQueue) StartQueue(ctx context.Context) {
	if p.isRunning {
		p.logger.Info(p.ctx, "Job already running")
		return
	}

	p.isRunning = true
	go p.startQueue()
	<-p.startReady
	p.logger.Info(p.ctx, "Fetch data job started")
}

func (p *domainEventQueue) startQueue() {
	close(p.startReady)

	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-p.ctx.Done():
			p.logger.Info(p.ctx, "Queue stopped via context: %v\n", p.ctx.Err())
			return
		case <-ticker.C:
			p.domainEventBus.Publish(p.ctx)
			p.logger.Info(p.ctx, "Queue working, time: ", time.Now().UTC())
		}
	}
}

func (p *domainEventQueue) StopQueue(ctx context.Context) {
	if p.isRunning {
		p.logger.Info(p.ctx, "Stopping fetch data job")
		p.cancel()
	}	
}