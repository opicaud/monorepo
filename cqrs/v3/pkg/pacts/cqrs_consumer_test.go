package pacts

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/opicaud/monorepo/grpc-eventstore/v2/inmemory/client"
	pkg "github.com/opicaud/monorepo/grpc-eventstore/v2/inmemory/client"
	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestSaveEvents(t *testing.T) {
	dir2, _ := filepath.Abs("../../../../grpc-eventstore/v2/proto/grpc_event_store.proto")
	grpcInteraction := `{
		"pact:proto": "` + dir2 + `",
		"pact:proto-service": "EventStore/Save",
		"pact:content-type": "application/protobuf",
		"request": {
			"event":
				{"aggregateId": {"id":"00000000-0000-0000-0000-000000000000"}, "name": "SHAPE_CREATED", "data": "e1wiTmF0dXJlXCI6XCJzcXVhcmVcIixcIkRpbWVuc2lvbnNcIjpbMiwzXSxcIklkXCI6XCIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDBcIixcIkFyZWFcIjoxfQ=="}
		},
		"response": {
			"status": "matching(number, 0)",
			"message": ""
		}
	}`

	var mockProvider, _ = message.NewSynchronousPact(message.Config{
		Consumer: "cqrs-save-events",
		Provider: "api-grpc-eventstore",
	})

	F := func(transport message.TransportConfig, m message.SynchronousMessage) error {
		from := pkg.NewInMemoryGrpcEventStoreFrom("127.0.0.1", transport.Port)
		parse, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
		_, _, err := from.Save(context.Background(), StandardEvent{aggregateId: parse,
			data: []byte("e1wiTmF0dXJlXCI6XCJzcXVhcmVcIixcIkRpbWVuc2lvbnNcIjpbMiwzXSxcIklkXCI6XCIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDBcIixcIkFyZWFcIjoxfQ=="),
			name: "SHAPE_CREATED"})
		return err
	}
	err := mockProvider.AddSynchronousMessage("save event").
		UsingPlugin(message.PluginConfig{
			Plugin:  "protobuf",
			Version: "0.3.5",
		}).
		WithContents(grpcInteraction, "application/grpc").
		StartTransport("grpc", "127.0.0.1", nil).
		ExecuteTest(t, F)

	assert.NoError(t, err)

}
