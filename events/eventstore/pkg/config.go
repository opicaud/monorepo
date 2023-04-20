package pkg

import (
	"fmt"
	pkg2 "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	"github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory"
	"github.com/opicaud/monorepo/events/pkg"
)

type Config interface {
	LoadConfig() (pkg.EventStore, error)
	SetDefaultConfig()
}

func (f *V1) LoadConfig() (pkg.EventStore, error) {
	return NewEventStoreBuilder().
		WithHost("localhost").
		WithPort(50051).
		Build(f.Protocol)
}

func (f *V1) SetDefaultConfig() {
	f.Protocol = "none"
}

type V1 struct {
	Protocol string
}

type Builder struct {
	host string
	port int
}

func (s *Builder) WithHost(host string) *Builder {
	s.host = host
	return s
}

func (s *Builder) WithPort(port int) *Builder {
	s.port = port
	return s
}

func (s *Builder) buildGrpc() (pkg.EventStore, error) {
	return pkg2.NewInMemoryGrpcEventStore(), nil
}

func (s *Builder) buildInMemory() (pkg.EventStore, error) {
	return inmemory.NewInMemoryEventStore(), nil
}

func (s *Builder) Build(protocol string) (pkg.EventStore, error) {
	switch protocol {
	case "none":
		return s.buildInMemory()
	case "grpc":
		return s.buildGrpc()
	default:
		return nil, fmt.Errorf("protocol %s not supported", protocol)
	}
}

func NewEventStoreBuilder() *Builder {
	s := new(Builder)
	return s
}
