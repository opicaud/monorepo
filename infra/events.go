package infra

import "github.com/google/uuid"

type EventStore interface {
	Save(events ...Event) error
	Load(id uuid.UUID) []Event
}

type EventsEmitter interface {
	NotifyAll(event ...Event)
	Add(subscriber Subscriber)
}

type Event interface {
	AggregateId() uuid.UUID
}

type Subscriber interface {
	Update(events []Event)
}
