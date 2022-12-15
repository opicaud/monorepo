package infra

import "github.com/google/uuid"

type eventStore interface {
	Save(events ...DomainEvent)
	Load(id uuid.UUID) ([]DomainEvent, error)
}

type eventsEmitter interface {
	NotifyAll(event ...DomainEvent)
	Add(subscriber Subscriber)
}

type DomainEvent interface {
	AggregateId() uuid.UUID
}

type Subscriber interface {
	Update(events []DomainEvent)
}
