package pkg

import (
	"fmt"
	"github.com/opicaud/monorepo/events/pkg"
	v2beta "github.com/opicaud/monorepo/events/pkg/v2beta"
	"github.com/spf13/viper"
	"log"
)

func NewEventsFrameworkFromConfig(path string) (pkg.EventStore, error) {
	config, errors := loadConfigFromPath(path, &V1{})
	return setConfig(config, errors).LoadConfig()
}

func NewEventsFrameworkFromConfigV2(path string) (v2beta.EventStore, error) {
	config, errors := loadConfigFromPathNew(path, &V2Beta{})
	return setConfig(config, errors).LoadConfig()
}

func setConfig(config Config, errors error) Config {
	if errors != nil {
		config.SetDefaultConfig()
		log.Printf("%s", errors)
		log.Printf("Loading default config due to previous errors..")
		log.Printf("Default config is %#v", config)
	}
	return config
}

func loadConfigFromPath(path string, config Config) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.UnmarshalKey("event-store", &config); err != nil {
		return config, err
	}
	return config, nil
}

func loadConfigFromPathNew(path string, config Config) (Config, error) {
	config, err := loadConfigFromPath(path, config)
	if err != nil {
		return config, err
	}
	if viper.GetString("version") != config.Version() {
		return config, fmt.Errorf("version not aligned between config file and your implementation")
	}
	return config, nil
}
