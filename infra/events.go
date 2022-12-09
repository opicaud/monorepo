package infra

import "github.com/google/uuid"

type EventStore interface {
	Save(events ...Event)
	Load(id uuid.UUID) ([]Event, error)
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
