package utils

import (
	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ContractTest struct {
	GrpcInteraction string
	F               func(tc message.TransportConfig, m message.SynchronousMessage) error
	Description     string
}

type ConsumerAndProvider struct {
	Consumer string
	Provider string
}

func newMockProvider(cp ConsumerAndProvider) (*message.SynchronousPact, error) {
	var mockProvider, err = message.NewSynchronousPact(message.Config{
		Consumer: cp.Consumer,
		Provider: cp.Provider,
	})

	return mockProvider, err
}

func NewMockProviderWithoutError(t *testing.T, cp ConsumerAndProvider) *message.SynchronousPact {
	var mockProvider, err = newMockProvider(cp)
	assert.NoError(t, err)
	return mockProvider
}

func ExecuteTest(t *testing.T, mockProvider *message.SynchronousPact, contractTest ContractTest) {
	err := mockProvider.
		AddSynchronousMessage(contractTest.Description).
		UsingPlugin(message.PluginConfig{
			Plugin: "protobuf",
		}).
		WithContents(contractTest.GrpcInteraction, "application/grpc").
		StartTransport("grpc", "127.0.0.1", nil).
		ExecuteTest(t, contractTest.F)

	assert.NoError(t, err)

}

func RunTest(t *testing.T, contractTest ContractTest, cp ConsumerAndProvider) {
	mockProvider := NewMockProviderWithoutError(t, cp)
	ExecuteTest(t, mockProvider, contractTest)
}
