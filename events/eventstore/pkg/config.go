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
	switch f.Protocol {
	case "none":
		return inmemory.NewInMemoryEventStore(), nil
	case "grpc":
		return pkg2.NewInMemoryGrpcEventStore(), nil
	default:
		return nil, fmt.Errorf("protocol %s not supported", f.Protocol)
	}
}

func (f *V1) SetDefaultConfig() {
	f.Protocol = "none"
}

type V1 struct {
	Protocol string
}
