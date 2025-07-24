package interfaces

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
)

type DomainEventBus interface{
	Publish() results.Result
}