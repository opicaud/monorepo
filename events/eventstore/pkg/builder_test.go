package pkg

import (
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolFromFile(t *testing.T) {
	_, err := NewEventsFrameworkFromConfig("internal/default_config.yml")
	assert.NoError(t, err)
}

func TestConfigProtocolFromFileV1(t *testing.T) {
	config, err := loadConfigFromPath("internal/v1.yml")
	assertTypeConfig(t, err, &V1{}, config)

}
func TestConfigProtocolFromFileV2(t *testing.T) {
	_, err := loadConfigFromPath("internal/v2.yml")
	assert.Error(t, err)
}

func TestConfigProtocolFromADummyFile(t *testing.T) {
	config, err := loadConfigFromPath("internal/not_ok_config.yml")
	assertTypeConfig(t, err, &V1{}, config)
}

func TestConfigProtocolDefaultConfig(t *testing.T) {
	noConfig := ""
	config, err := loadConfigFromPath(noConfig)
	assert.NoError(t, err)
	assertTypeConfig(t, err, &V1{}, config)
}

func assertType(t *testing.T, err error, expected pkg.EventStore, actual pkg.EventStore) {
	assert.NoError(t, err)
	assert.IsType(t, expected, actual)
}

func assertTypeConfig(t *testing.T, err error, expected Config, actual Config) {
	assert.NoError(t, err)
	assert.IsType(t, expected, actual)
}
