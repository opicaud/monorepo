package pkg

import (
	"bytes"
	"fmt"
	pkg2 "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	"github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolNone(t *testing.T) {
	provider, err := newEventsFrameworkBuilderFromConfig(protocol("none"))
	assertType(t, err, &inmemory.EventStore{}, provider)
}

func TestConfigProtocolGrpc(t *testing.T) {
	provider, err := newEventsFrameworkBuilderFromConfig(protocol("grpc"))
	assertType(t, err, &pkg2.InMemoryGrpcEventStore{}, provider)
}

func TestConfigProtocolFromFile(t *testing.T) {
	_, err := NewEventsFrameworkFromConfig("internal/test_config.yml")
	assert.NoError(t, err)
}

func TestConfigProtocolDefaultConfig(t *testing.T) {
	provider, err := NewEventsFrameworkFromConfig("")
	assertType(t, err, &inmemory.EventStore{}, provider)

}

func assertType(t *testing.T, err error, expected pkg.EventStore, actual pkg.EventStore) {
	assert.NoError(t, err)
	assert.IsType(t, expected, actual)

}

func protocol(s string) *bytes.Buffer {
	var yamlExample = []byte(`
event-store:
  protocol: ` + s + `
`)
	return bytes.NewBuffer(yamlExample)
}

func newEventsFrameworkBuilderFromConfig(s *bytes.Buffer) (pkg.EventStore, error) {
	viper.SetConfigType("yaml")
	loadConfigForTest(s)
	return loadConfigV1()
}

func loadConfigForTest(s *bytes.Buffer) {
	err := viper.ReadConfig(s) // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
