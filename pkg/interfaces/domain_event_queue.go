package interfaces

import "context"

type DomainEventQueue interface {
	StartQueue(ctx context.Context)
	StopQueue(ctx context.Context)
}