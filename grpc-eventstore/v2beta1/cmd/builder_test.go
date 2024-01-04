package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolFromFile(t *testing.T) {
	eventStore, err := NewEventsFrameworkFromConfig("../internal/v2beta1.yml")
	assert.NoError(t, err)
	assert.IsType(t, &InMemoryEventStoreWithoutGrpc{}, eventStore)
}
