package pkg

import (
	pkg "github.com/opicaud/monorepo/cqrs/v3/pkg"
	pkg2 "github.com/opicaud/monorepo/grpc-eventstore/v2/inmemory/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolNoneV2(t *testing.T) {
	testProtocol(t, "none", &InMemoryEventStoreWithoutGrpc{})
}

func TestConfigProtocolGrpcV2(t *testing.T) {
	testProtocol(t, "grpc", &pkg2.InMemoryGrpcEventStore{})
}

func testProtocol(t *testing.T, protocol string, expectedType pkg.EventStore) {
	eventStore, err := loadV2Config(protocol)
	assert.NoError(t, err)
	assert.IsType(t, expectedType, eventStore)
}

func loadV2Config(protocol string) (pkg.EventStore, error) {
	v2 := V2{Protocol: protocol}
	eventStore, err := v2.LoadConfig()
	return eventStore, err
}
