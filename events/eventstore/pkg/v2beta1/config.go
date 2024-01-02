package pkg

import (
	"fmt"
	inmemory "github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory/v2beta1"
)

type Config interface {
	LoadConfig() (*inmemory.EventStore, error)
	SetDefaultConfig()
	Version() string
}

type V2Beta1 struct {
	Protocol string
	Port     int
	Host     string
}

func (f *V2Beta1) LoadConfig() (*inmemory.EventStore, error) {
	return NewEventStoreBuilder().
		WithHost(f.Host).
		WithPort(f.Port).
		Build(f.Protocol)
}

func (f *V2Beta1) SetDefaultConfig() {
	f.Protocol = "none"
}

func (f *V2Beta1) Version() string {
	return "v2/beta1"
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

func (s *Builder) Build(protocol string) (*inmemory.EventStore, error) {
	switch protocol {
	case "none":
		return inmemory.NewInMemoryEventStore(), nil
	default:
		return nil, fmt.Errorf("protocol %s not supported", protocol)
	}
}

func NewEventStoreBuilder() *Builder {
	s := new(Builder)
	return s
}