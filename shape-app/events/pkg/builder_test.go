package pkg

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAStandardHandlerACommand(t *testing.T) {
	fake := &fakeEventStore{}
	provider := NewEventsFrameworkBuilder().
		WithEventStore(fake).
		Build()
	assert.IsType(t, fake, provider.eventstore)
}

type fakeEventStore struct {
}

func (f *fakeEventStore) Save(events ...DomainEvent) error {
	return nil
}

func (f fakeEventStore) Load(uuid uuid.UUID) ([]DomainEvent, error) {
	return nil, nil
}
