package pkg

import (
	pkg "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg/v2beta1"
	pkg2 "github.com/opicaud/monorepo/events/pkg/v2beta1"

	inmemory "github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory/v2beta1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolNoneV2Beta1(t *testing.T) {
	testProtocol(t, "none", &inmemory.EventStore{})
}

func TestConfigProtocolGrpcV2Beta1(t *testing.T) {
	testProtocol(t, "grpc", &pkg.InMemoryGrpcEventStore{})
}

func testProtocol(t *testing.T, protocol string, expectedType pkg2.EventStore) {
	eventStore, err := loadV2Beta1Config(protocol)
	assert.NoError(t, err)
	assert.IsType(t, expectedType, eventStore)
}

func loadV2Beta1Config(protocol string) (pkg2.EventStore, error) {
	v2beta1 := V2Beta1{Protocol: protocol}
	eventStore, err := v2beta1.LoadConfig()
	return eventStore, err
}
