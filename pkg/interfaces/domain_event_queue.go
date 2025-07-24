package interfaces

type DomainEventQueue interface{
	StartQueue()
	StopQueue()
}