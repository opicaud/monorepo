package pkg

import (
	"fmt"
	pkg2 "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	"github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory"
	v2 "github.com/opicaud/monorepo/events/pkg/v2"
)

type Config interface {
	LoadConfig() (v2.EventStore, error)
	SetDefaultConfig()
}

func (f *V1) LoadConfig() (v2.EventStore, error) {
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

func (f *V2Beta) LoadConfig() (v2.EventStore, error) {
	return NewEventStoreBuilder().
		WithHost("localhost").
		WithPort(50051).
		Build(f.Protocol)
}

func (f *V2Beta) SetDefaultConfig() {
	f.Protocol = "none"
}

type V2Beta struct {
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

func (s *Builder) Build(protocol string) (v2.EventStore, error) {
	switch protocol {
	case "none":
		return inmemory.NewInMemoryEventStore(), nil
	case "grpc":
		return pkg2.NewInMemoryGrpcEventStore(), nil
	default:
		return nil, fmt.Errorf("protocol %s not supported", protocol)
	}
}

func NewEventStoreBuilder() *Builder {
	s := new(Builder)
	return s
}
