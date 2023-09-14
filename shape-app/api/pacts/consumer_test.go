package pacts

import (
	"context"
	"fmt"
	ac "github.com/opicaud/monorepo/shape-app/api/proto"
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

func TestCreateShape2(t *testing.T) {
	err2 := os.Getenv("PACT_PLUGIN_DIR")
	if err2 == "" {
		return
	}
	dir2, _ := filepath.Abs("../proto/app_shape.proto")
	grpcInteraction := `{
		"pact:proto": "` + dir2 + `",
		"pact:proto-service": "Shapes/create",
		"pact:content-type": "application/protobuf",
		"request": {
			"shapes": 
				{
					"shape": "rectangle",
					"dimensions": [ "matching(number, 3)", "matching(number, 4)"]
				}
			
		},
		"response": {
			"message": { "code": "matching(number, 0)"}
		}
	}`
	F := func(transport message.TransportConfig, m message.SynchronousMessage) error {
		dimensions := []float32{3, 4}
		shape := "rectangle"
		rectangle := ac.ShapeMessage{Shape: shape, Dimensions: dimensions}
		request := &ac.ShapeRequest{Shapes: &rectangle}
		area, err := getRectangleAndSquareArea2(fmt.Sprintf("localhost:%d", transport.Port), request)

		assert.NoError(t, err)
		assert.Equal(t, uint32(0), area.GetCode())

		return err
	}

	var mockProvider, _ = message.NewSynchronousPact(message.Config{
		Consumer: "grpc-consumer-go",
		Provider: "area-calculator-provider",
	})

	_ = mockProvider.AddSynchronousMessage("calculate rectangle area request").
		UsingPlugin(message.PluginConfig{
			Plugin:  "protobuf",
			Version: "0.3.5",
		}).
		WithContents(grpcInteraction, "application/grpc").
		StartTransport("grpc", "127.0.0.1", nil).
		ExecuteTest(t, F)

}

func getRectangleAndSquareArea2(address string, request *ac.ShapeRequest) (*ac.Message, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := ac.NewShapesClient(conn)

	log.Println("Sending calculate rectangle and square request")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Create(ctx, request)

	if err != nil {
		return nil, err
	}
	return r.GetMessage(), nil
}
