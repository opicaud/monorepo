package eventstore

import "github.com/google/uuid"

type EventStore interface {
	Save(events ...DomainEvent) error
	Load(id uuid.UUID) ([]DomainEvent, error)
}

type EventsEmitter interface {
	NotifyAll(event ...DomainEvent)
	Add(subscriber Subscriber)
}

type DomainEvent interface {
	AggregateId() uuid.UUID
}

type Subscriber interface {
	Update(events []DomainEvent)
}
