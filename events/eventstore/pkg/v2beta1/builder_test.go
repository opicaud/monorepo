package pkg

import (
	inmemory "github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory/v2beta1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolFromFile(t *testing.T) {
	eventStore, err := NewEventsFrameworkFromConfig("../internal/v2beta1.yml")
	assert.NoError(t, err)
	assert.IsType(t, &inmemory.EventStore{}, eventStore)
}
