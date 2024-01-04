package pkg

import (
	"fmt"
	pkg "github.com/opicaud/monorepo/cqrs/pkg/v3beta1"
	"github.com/spf13/viper"
	"log"
)

func NewEventsFrameworkFromConfig(path string) (pkg.EventStore, error) {
	config, errors := loadConfigFromPathNew(path, &V2Beta1{})
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
