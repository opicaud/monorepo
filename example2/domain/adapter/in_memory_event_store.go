package adapter

import (
	"fmt"
	"github.com/google/uuid"
)

func NewInMemoryEventStore() *InMemoryEventStore {
	fakeRepository := new(InMemoryEventStore)
	return fakeRepository
}

func (f *InMemoryEventStore) Save(events ...DomainEvent) error {
	f.events = append(f.events, events...)
	return nil
}

func (f InMemoryEventStore) Load(uuid uuid.UUID) ([]DomainEvent, error) {
	w := 0
	for _, e := range f.events {
		if e.AggregateId() == uuid {
			f.events[w] = e
			w++
		}
	}
	if len(f.events[0:w]) == 0 {
		return nil, fmt.Errorf("No aggregate with id %s has been found", uuid)
	}
	return f.events[0:w], nil
}

type InMemoryEventStore struct {
	events []DomainEvent
}
