package pacts

import (
	"context"
	"fmt"
	"github.com/bazelbuild/rules_go/go/runfiles"
	gen "github.com/opicaud/monorepo/events/eventstore/grpc/proto"
	pact "github.com/opicaud/monorepo/pact-helper/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"
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
			"id": "a-uuid"
		},
		"response": {
			"status": "matching(number, 0)"
		}
	}`
	log.Println(grpcInteraction)
	var c = new(pact.ContractTest)
	c.GrpcInteraction = grpcInteraction
	c.Description = "load events from eventstore"
	c.F = func(transport message.TransportConfig, m message.SynchronousMessage) error {
		request := &gen.UUID{Id: "a-uuid"}
		_, err := loadEvents(fmt.Sprintf("localhost:%d", transport.Port), request)

		assert.NoError(t, err)
		return err
	}

	pact.RunTest(t, *c, pact.ConsumerAndProvider{
		Consumer: "area-calculator-grpc-api",
		Provider: "grpc-eventstore",
	})

}

func loadEvents(address string, request *gen.UUID) (*gen.Response, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := gen.NewEventStoreClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Load(ctx, request)

	if err != nil {
		panic("Error has occured, stop")
	}
	return r, nil
}
