package pkg

import (
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/spf13/viper"
	"log"
)

func NewEventsFrameworkFromConfig(path string) (pkg.EventStore, error) {
	config, errors := loadConfigFromPath(path)
	if errors != nil {
		config.SetDefaultConfig()
		log.Printf("%s", errors)
		log.Printf("Loading default config due to previous errors..")
	}
	return config.LoadConfig()
}

func loadConfigFromPath(path string) (Config, error) {
	viper.SetConfigFile(path)
	var config = fetchConfigVersion()
	if err := viper.ReadInConfig(); err != nil {
		return &V1{}, err
	}
	if err := viper.UnmarshalKey("event-store", &config); err != nil {
		return &V1{}, err
	}
	return config, nil
}

func fetchConfigVersion() Config {
	switch viper.GetString("version") {
	case "v1":
		return &V1{}
	default:
		log.Println("Version not found in config, load by default version:v1")
		return &V1{Protocol: "none"}

	}
}
