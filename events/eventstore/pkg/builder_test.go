package pkg

import (
	"bytes"
	"fmt"
	pkg2 "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	"github.com/opicaud/monorepo/events/eventstore/inmemory/cmd"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolNone(t *testing.T) {
	provider, err := newEventsFrameworkBuilderFromConfig(protocol("none"))
	assertType(t, err, &cmd.InMemoryEventStore{}, provider)
}

func TestConfigProtocolGrpc(t *testing.T) {
	provider, err := newEventsFrameworkBuilderFromConfig(protocol("grpc"))
	assertType(t, err, &pkg2.InMemoryGrpcEventStore{}, provider)
}

func TestConfigProtocolFromFile(t *testing.T) {
	_, err := NewEventsFrameworkFromConfig("internal/test_config.yml")
	assert.NoError(t, err)
}

func assertType(t *testing.T, err error, expected pkg.EventStore, provider *EventStoreProvider) {
	assert.NoError(t, err)
	assert.IsType(t, expected, provider.eventStore)

}

func protocol(s string) *bytes.Buffer {
	var yamlExample = []byte(`
event-store:
  protocol: ` + s + `
`)
	return bytes.NewBuffer(yamlExample)
}

func newEventsFrameworkBuilderFromConfig(s *bytes.Buffer) (*EventStoreProvider, error) {
	viper.SetConfigType("yaml")
	loadConfig(s)
	return loadProtocol()
}

func loadConfig(s *bytes.Buffer) {
	err := viper.ReadConfig(s) // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
