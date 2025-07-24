package enums

import "fmt"


type EventType int

const (
	EventDeleted EventType = iota
	UserWantsToVisitEvent
	UserCanVisitEvent
	UserDeleted
)

func (c EventType) String() string {
	switch c {
	case 0:
		return "EventDeleted"
	case 1:
		return "UserWantsToVisitEvent"
	case 2:
		return "UserCanVisitEvent"
	case 3:
		return "UserDeleted"
	}
	return fmt.Sprintf("(%q)", int(c))
}