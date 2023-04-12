package pkg

import (
	"bytes"
	"fmt"
	grpc "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	"github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/spf13/viper"
	"log"
)

func loadConfig() (pkg.EventStore, error) {
	protocol := viper.GetString("event-store.protocol")
	log.Printf("Loading protocol: %s\n", protocol)
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
	err := viper.ReadInConfig()
	if err != nil {
		return setDefaultConfig()
	}
	return loadConfig()

}

func setDefaultConfig() (pkg.EventStore, error) {
	log.Println("Loading default protocol..")
	viper.SetConfigType("yaml")
	_ = viper.ReadConfig(defaultConfigYaml())
	return loadConfig()
}

func defaultConfigYaml() *bytes.Buffer {
	var defaultConfig = []byte(`
event-store:
  protocol: none
`)
	return bytes.NewBuffer(defaultConfig)
}
