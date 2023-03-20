package pkg

import (
	"fmt"
	"github.com/google/uuid"
	pkg2 "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	"github.com/opicaud/monorepo/events/eventstore/inmemory/cmd"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/spf13/viper"
)

func newEventsFrameworkBuilder() *Builder {
	return &Builder{}
}

func (s *Builder) withEventStore(eventStore pkg.EventStore) *Builder {
	s.eventStore = eventStore
	return s
}

func (s *Builder) build() *EventStoreProvider {
	infra := new(EventStoreProvider)
	infra.eventStore = s.eventStore
	return infra
}

type Builder struct {
	eventStore pkg.EventStore
}

type EventStoreProvider struct {
	eventStore pkg.EventStore
}

func (f *EventStoreProvider) Save(events ...pkg.DomainEvent) {
	err := f.eventStore.Save(events...)
	if err != nil {
		err = fmt.Errorf("error has occured when save events")
		fmt.Println(err.Error())
	}
}

func (f *EventStoreProvider) Load(uuid uuid.UUID) ([]pkg.DomainEvent, error) {
	return f.eventStore.Load(uuid)
}

func loadProtocol() (*EventStoreProvider, error) {
	protocol := viper.GetString("event-store.protocol")
	builder := newEventsFrameworkBuilder()
	switch protocol {
	case "none":
		builder.
			withEventStore(cmd.NewInMemoryEventStore())
	case "grpc":
		builder.
			withEventStore(pkg2.NewInMemoryGrpcEventStore())
	default:
		return nil, fmt.Errorf("protocol %s not supported", protocol)
	}
	return builder.build(), nil
}

func NewEventsFrameworkFromConfig(s string) (*EventStoreProvider, error) {
	viper.SetConfigFile(s)
	_ = viper.ReadInConfig()
	return loadProtocol()

}
