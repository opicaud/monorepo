package pkg

import (
	"github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolFromFile(t *testing.T) {
	_, err := NewEventsFrameworkFromConfig("internal/default_config.yml")
	assert.NoError(t, err)
}

func TestConfigProtocolFromFileV1(t *testing.T) {
	_, err := NewEventsFrameworkFromConfig("internal/v1.yml")
	assert.NoError(t, err)
}
func TestConfigProtocolFromFileV2(t *testing.T) {
	_, err := NewEventsFrameworkFromConfig("internal/v2.yml")
	assert.NoError(t, err)
}

func TestConfigProtocolFromADummyFile(t *testing.T) {
	_, err := NewEventsFrameworkFromConfig("internal/not_ok_config.yml")
	assert.Error(t, err)
}

func TestConfigProtocolDefaultConfig(t *testing.T) {
	noConfig := ""
	provider, err := NewEventsFrameworkFromConfig(noConfig)
	assert.NoError(t, err)
	assertType(t, err, &inmemory.EventStore{}, provider)
}

func assertType(t *testing.T, err error, expected pkg.EventStore, actual pkg.EventStore) {
	assert.NoError(t, err)
	assert.IsType(t, expected, actual)
}
