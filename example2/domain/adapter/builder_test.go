package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAStandardHandlerACommand(t *testing.T) {
	provider := NewInfraBuilder().
		WithEventStore(NewInMemoryEventStore()).
		Build()
	assert.IsType(t, &InMemoryEventStore{}, provider.eventstore)
}
