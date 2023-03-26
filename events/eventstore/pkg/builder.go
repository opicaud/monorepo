package pkg

import (
	"fmt"
	grpc "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	"github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/spf13/viper"
)

func loadProtocol() (pkg.EventStore, error) {
	protocol := viper.GetString("event-store.protocol")
	switch protocol {
	case "none":
		return inmemory.NewInMemoryEventStore(), nil
	case "grpc":
		return grpc.NewInMemoryGrpcEventStore(), nil
	default:
		return nil, fmt.Errorf("protocol %s not supported", protocol)
	}

}

func NewEventsFrameworkFromConfig(s string) (pkg.EventStore, error) {
	viper.SetConfigFile(s)
	_ = viper.ReadInConfig()
	return loadProtocol()

}
