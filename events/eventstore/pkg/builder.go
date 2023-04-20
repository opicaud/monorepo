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
	if err := viper.ReadInConfig(); err != nil {
		return &V1{}, err
	}
	var config = fetchConfigVersion()
	if err := viper.UnmarshalKey("event-store", &config); err != nil {
		return &V1{}, err
	}
	return config, nil
}

func fetchConfigVersion() Config {
	version := viper.GetString("version")
	log.Printf("version in config file: %s", version)
	switch version {
	case "v1":
		return &V1{}
	case "v2/beta":
		return &V2Beta{}
	default:
		log.Println("version not found in config, load by default v1")
		return &V1{Protocol: "none"}
	}
}
