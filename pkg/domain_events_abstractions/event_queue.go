package domain_events_abstractions

type EventQueue interface{
	StartQueue()
	StopQueue()
}