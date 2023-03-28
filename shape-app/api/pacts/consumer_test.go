package pacts

import (
	"context"
	"fmt"
	"github.com/bazelbuild/rules_go/go/runfiles"
	pact "github.com/opicaud/monorepo/pact-helper/go"
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

var cp = &pact.ConsumerAndProvider{
	Consumer: "grpc-consumer-go",
	Provider: "area-calculator-provider",
}

var dir, _ = filepath.Abs("../proto/app_shape.proto")

func TestCreateShape(t *testing.T) {
	rlocation, err2 := runfiles.Rlocation(os.Getenv("SHAPEAPP_PROTO_FILE"))
	if err2 != nil {
		rlocation = dir
	}
	grpcInteraction := `{
		"pact:proto": "` + rlocation + `",
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

	var c = new(pact.ContractTest)
	c.GrpcInteraction = grpcInteraction
	c.Description = "calculate rectangle area request"
	c.F = func(transport message.TransportConfig, m message.SynchronousMessage) error {
		dimensions := []float32{3, 4}
		shape := "rectangle"
		rectangle := ac.ShapeMessage{Shape: shape, Dimensions: dimensions}
		request := &ac.ShapeRequest{Shapes: &rectangle}
		area, err := getRectangleAndSquareArea(fmt.Sprintf("localhost:%d", transport.Port), request)

		assert.NoError(t, err)
		assert.Equal(t, uint32(0), area.GetCode())

		return err
	}

	pact.RunTest(t, *c, *cp)

}

func getRectangleAndSquareArea(address string, request *ac.ShapeRequest) (*ac.Message, error) {
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
