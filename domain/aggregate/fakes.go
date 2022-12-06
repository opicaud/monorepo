package aggregate

import (
	"example2/infra"
)

func NewInMemoryEventStore() *InMemoryRepository {
	fakeRepository := new(InMemoryRepository)
	return fakeRepository
}

func (f *InMemoryRepository) Save(events ...infra.Event) error {
	return nil
}

type InMemoryRepository struct {
	Shapes []infra.Event
}
