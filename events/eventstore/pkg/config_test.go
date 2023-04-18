package pkg

import (
	"github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	"github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory"
	v1 "github.com/opicaud/monorepo/events/pkg"
	"testing"
)

func TestConfigProtocolNone(t *testing.T) {
	eventStore, err := loadV1Config("none")
	assertType(t, err, &inmemory.EventStore{}, eventStore)
}

func TestConfigProtocolGrpc(t *testing.T) {
	eventStore, err := loadV1Config("grpc")
	assertType(t, err, &pkg.InMemoryGrpcEventStore{}, eventStore)

}

func loadV1Config(protocol string) (v1.EventStore, error) {
	v1 := V1{Protocol: protocol}
	eventStore, err := v1.LoadConfig()
	return eventStore, err
}
