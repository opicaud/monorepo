package pkg

import (
	"fmt"
	pkg2 "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	"github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory"
	v2beta "github.com/opicaud/monorepo/events/pkg/v2beta"
)

type Config interface {
	LoadConfig() (v2beta.EventStore, error)
	SetDefaultConfig()
	Version() string
}

func (f *V1) LoadConfig() (v2beta.EventStore, error) {
	return NewEventStoreBuilder().
		WithHost("localhost").
		WithPort(50051).
		Build(f.Protocol)
}

func (f *V1) SetDefaultConfig() {
	f.Protocol = "none"
}

func (f *V1) Version() string {
	return "v1"
}

type V1 struct {
	Protocol string
}

func (f *V2Beta) LoadConfig() (v2beta.EventStore, error) {
	return NewEventStoreBuilder().
		WithHost(f.Host).
		WithPort(f.Port).
		Build(f.Protocol)
}

func (f *V2Beta) SetDefaultConfig() {
	f.Protocol = "none"
}

func (f *V2Beta) Version() string {
	return "v2/beta"
}

type V2Beta struct {
	Protocol string
	Port     int
	Host     string
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

func (s *Builder) Build(protocol string) (v2beta.EventStore, error) {
	switch protocol {
	case "none":
		return inmemory.NewInMemoryEventStore(), nil
	case "grpc":
		return pkg2.NewInMemoryGrpcEventStoreFrom(s.host, s.port), nil
	default:
		return nil, fmt.Errorf("protocol %s not supported", protocol)
	}
}

func NewEventStoreBuilder() *Builder {
	s := new(Builder)
	return s
}
