package pkg

import (
	pkg "github.com/opicaud/monorepo/cqrs/pkg/v3beta1"
	pkg2 "github.com/opicaud/monorepo/grpc-eventstore/v2beta1/inmemory/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolNoneV2Beta1(t *testing.T) {
	testProtocol(t, "none", &InMemoryEventStoreWithoutGrpc{})
}

func TestConfigProtocolGrpcV2Beta1(t *testing.T) {
	testProtocol(t, "grpc", &pkg2.InMemoryGrpcEventStore{})
}

func testProtocol(t *testing.T, protocol string, expectedType pkg.EventStore) {
	eventStore, err := loadV2Beta1Config(protocol)
	assert.NoError(t, err)
	assert.IsType(t, expectedType, eventStore)
}

func loadV2Beta1Config(protocol string) (pkg.EventStore, error) {
	v2beta1 := V2Beta1{Protocol: protocol}
	eventStore, err := v2beta1.LoadConfig()
	return eventStore, err
}
