package pkg

import (
	"github.com/opicaud/monorepo/events/eventstore/pkg/internal/inmemory/v2beta1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigProtocolNoneV2Beta1(t *testing.T) {
	eventStore, err := loadV2Beta1Config("none")
	assert.NoError(t, err)
	assert.IsType(t, &inmemory.EventStore{}, eventStore)
}

func loadV2Beta1Config(protocol string) (*inmemory.EventStore, error) {
	v2beta1 := V2Beta1{Protocol: protocol}
	eventStore, err := v2beta1.LoadConfig()
	return eventStore, err
}
