package infra

type EventBus interface {
	NotifyAll()
}

type EventsEmitter interface {
	DispatchEvent(event ...Event)
}

type StandardEventsEmitter struct{}

func (s *StandardEventsEmitter) DispatchEvent(event ...Event) {}

type Event interface{}
