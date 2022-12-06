package infra

import (
	"github.com/google/uuid"
)

func NewInMemoryEventStore() *InMemoryEventStore {
	fakeRepository := new(InMemoryEventStore)
	return fakeRepository
}

func (f *InMemoryEventStore) Save(events ...Event) error {
	f.events = append(f.events, events...)
	return nil
}

func (f InMemoryEventStore) Load(uuid uuid.UUID) []Event {
	w := 0
	for _, e := range f.events {
		if e.AggregateId() == uuid {
			f.events[w] = e
			w++
		}
	}
	return f.events[0:w]
}

type InMemoryEventStore struct {
	events []Event
}
