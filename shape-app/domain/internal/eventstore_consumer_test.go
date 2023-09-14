package internal

import (
	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	gen "github.com/opicaud/monorepo/events/eventstore/grpc/proto"
	pkg2 "github.com/opicaud/monorepo/events/pkg"
	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadEvents(t *testing.T) {
	err2 := os.Getenv("PACT_PLUGIN_DIR")
	if err2 == "" {
		return
	}
	dir2, _ := filepath.Abs("../../../events/eventstore/grpc/proto/grpc_event_store.proto")
	grpcInteraction := `{
		"pact:proto": "` + dir2 + `",
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
	//SetEnvVarPactPluginDir()

	_ = mockProvider.AddSynchronousMessage("fetch event").GivenWithParameter(models.ProviderState{
		Name: "a state",
		Parameters: map[string]interface{}{
			"events": `[{"aggregateId": {"id":"00000000-0000-0000-0000-000000000000"}, "name": "SHAPE_CREATED", "data": "eyJOYXR1cmUiOiJzcXVhcmUiLCJEaW1lbnNpb25zIjpbMiwzXSwiSWQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJBcmVhIjoxfQ=="}]`,
		},
	}).
		UsingPlugin(message.PluginConfig{
			Plugin:  "protobuf",
			Version: "0.3.5",
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

	zipPath := os.Getenv("PACT_PLUGINS_ZIP")
	zipFileLocation, _ := r.Rlocation(zipPath)
	pluginDir := unzip(zipFileLocation)
	_ = os.Setenv("PACT_PLUGIN_DIR", pluginDir)
	log.Printf("PACT_PLUGIN_DIR: %s", pluginDir)

	if err != nil {
		log.Fatalf("path %s not found", zipPath)
	}

}

func unzip(zipFile string) string {
	version := module.Version{
		Path:    "pact.plugins.protobuf",
		Version: "v0.3.6",
	}

	if _, err := os.Stat("./protobuf-0.3.6"); os.IsNotExist(err) {
		log.Println("Unzipping plugins..")
		uz := zip.Unzip("protobuf-0.3.6", version, zipFile)

		if uz != nil {
			log.Panic(uz.Error())
		}
		err := os.Chmod("./protobuf-0.3.6/pact-protobuf-plugin", 0777)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("Plugins already present, skipping unzip process..")
	}
	getwd, _ := os.Getwd()
	return getwd

}
