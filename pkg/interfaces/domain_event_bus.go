package interfaces

import (
	"context"

	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
)

type DomainEventBus interface{
	Publish(ctx context.Context) results.Result
}