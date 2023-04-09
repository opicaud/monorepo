package internal

import (
	"fmt"
	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	gen "github.com/opicaud/monorepo/events/eventstore/grpc/proto"
	pkg2 "github.com/opicaud/monorepo/events/pkg"
	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func GetProtoDir() string {
	dir, err := filepath.Abs("../../../events/eventstore/grpc/proto/grpc_event_store.proto")
	if err != nil {
		panic(fmt.Sprintf("File not found; %s", dir))
	}
	return dir
}
func TestLoadEvents(t *testing.T) {
	rlocation, err2 := runfiles.Rlocation(os.Getenv("EVENTSTORE_PROTO_FILE"))
	if err2 != nil {
		rlocation = GetProtoDir()
	}
	grpcInteraction := `{
		"pact:proto": "` + rlocation + `",
		"pact:proto-service": "EventStore/Load",
		"pact:content-type": "application/protobuf",
		"request": {
			"id": "00000000-0000-0000-0000-000000000000"
		},
		"response": {
			"status": "matching(number, 0)",
			"events":
				{
					"event": [{"aggregateId": {"id":"00000000-0000-0000-0000-000000000000"}, "name": "SHAPE_CREATED", "data": "{\"Nature\":\"square\",\"Dimensions\":[2,3],\"Id\":\"00000000-0000-0000-0000-000000000000\",\"Area\":1}"}]
				},
			"message": ""
		}
	}`

	var mockProvider, _ = message.NewSynchronousPact(message.Config{
		Consumer: "area-calculator-grpc",
		Provider: "api-grpc-eventstore",
	})

	F := func(transport message.TransportConfig, m message.SynchronousMessage) error {
		request := &gen.UUID{Id: "00000000-0000-0000-0000-000000000000"}
		id, _ := uuid.Parse(request.GetId())
		events, err := loadEvents("localhost", transport.Port, id)
		assert.Len(t, events, 1)
		event := NewShapeEventFactory().NewDeserializedEvent(id, events[0])
		assert.IsType(t, &Created{}, event)
		assert.Equal(t, "SHAPE_CREATED", event.Name())
		assert.Equal(t, id, event.AggregateId())
		assert.Equal(t, "square", event.(*Created).Nature)
		assert.Equal(t, []float32{2, 3}, event.(*Created).Dimensions)
		assert.Equal(t, float32(1), event.(*Created).Area)

		if err != nil {
			return err
		}
		return nil
	}
	SetEnvVarPactPluginDir()

	_ = mockProvider.AddSynchronousMessage("fetch event").GivenWithParameter(models.ProviderState{
		Name: "a state",
		Parameters: map[string]interface{}{
			"events": `[{"aggregateId": {"id":"00000000-0000-0000-0000-000000000000"}, "name": "SHAPE_CREATED", "data": "eyJOYXR1cmUiOiJzcXVhcmUiLCJEaW1lbnNpb25zIjpbMiwzXSwiSWQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJBcmVhIjoxfQ=="}]`,
		},
	}).
		UsingPlugin(message.PluginConfig{
			Plugin: "protobuf",
		}).
		WithContents(grpcInteraction, "application/grpc").
		StartTransport("grpc", "127.0.0.1", nil).
		ExecuteTest(t, F)

}

func loadEvents(address string, port int, id uuid.UUID) ([]pkg2.DomainEvent, error) {
	from := pkg.NewInMemoryGrpcEventStoreFrom(address, port)
	return from.Load(id)
}

func SetEnvVarPactPluginDir() {
	r, err := runfiles.New()
	if err != nil {
		log.Printf("Bazel not present, use PACT_PLUGIN_DIR: %s\n", os.Getenv("PACT_PLUGIN_DIR"))
		return
	}

	path := os.Getenv("PACT_PLUGINS")
	pactPluginDr, err := r.Rlocation(path)
	_ = os.Setenv("PACT_PLUGIN_DIR", filepath.Dir(pactPluginDr))

	log.Printf("PACT_PLUGIN_DIR: %s", filepath.Dir(pactPluginDr))
	if err != nil {
		log.Fatalf("path %s not found", path)
	}

}
