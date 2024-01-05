package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolFromFile(t *testing.T) {
	eventStore, err := NewEventsFrameworkFromConfig("./internal/v2.yml")
	assert.NoError(t, err)
	assert.IsType(t, &InMemoryEventStoreWithoutGrpc{}, eventStore)
}

func TestConfigWithVersionNotAligned(t *testing.T) {
	eventStore, err := NewEventsFrameworkFromConfig("./internal/vX.yml")
	assert.NoError(t, err)
	assert.IsType(t, &InMemoryEventStoreWithoutGrpc{}, eventStore)
}
