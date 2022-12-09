package infra

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func NewInMemoryEventStore() *InMemoryEventStore {
	fakeRepository := new(InMemoryEventStore)
	return fakeRepository
}

func (f *InMemoryEventStore) Save(events ...Event) {
	f.events = append(f.events, events...)
}

func (f InMemoryEventStore) Load(uuid uuid.UUID) ([]Event, error) {
	w := 0
	for _, e := range f.events {
		if e.AggregateId() == uuid {
			f.events[w] = e
			w++
		}
	}
	if len(f.events[0:w]) == 0 {
		return nil, errors.New(fmt.Sprintf("No aggregate with id %s has been found", uuid))
	}
	return f.events[0:w], nil
}

type InMemoryEventStore struct {
	events []Event
}
