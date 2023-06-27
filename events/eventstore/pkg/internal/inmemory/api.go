package inmemory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewInMemoryEventStore() *EventStore {
	fakeRepository := new(EventStore)
	return fakeRepository
}

func (f *EventStore) Save(events ...pkg.DomainEvent) error {
	f.events = append(f.events, events...)
	return nil
}

func (f *EventStore) Remove(uuid uuid.UUID) error {
	f.events = []pkg.DomainEvent{}
	return nil
}

func (f *EventStore) Load(uuid uuid.UUID) ([]pkg.DomainEvent, error) {
	w := 0
	for _, e := range f.events {
		if e.AggregateId() == uuid {
			f.events[w] = e
			w++
		}
	}
	if len(f.events[0:w]) == 0 {
		return nil, fmt.Errorf("No aggregate with Id %s has been found", uuid)
	}
	return f.events[0:w], nil
}

type EventStore struct {
	events []pkg.DomainEvent
}

func (f *EventStore) GetHealthClient() grpc_health_v1.HealthClient {
	panic("GetHealthClient not implemented")
}
