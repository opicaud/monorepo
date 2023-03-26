package pacts

import (
	"fmt"
	"github.com/bazelbuild/rules_go/go/runfiles"
	pact "github.com/opicaud/monorepo/pact-helper/go"
	ac "github.com/opicaud/monorepo/shape-app/api/proto"
	"os"
	"testing"

	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"
)

var cp = &pact.ConsumerAndProvider{
	Consumer: "grpc-consumer-go",
	Provider: "area-calculator-provider",
}

func TestCreateShape(t *testing.T) {
	rlocation, _ := runfiles.Rlocation(os.Getenv("SHAPEAPP_PROTO_FILE"))
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
		area, err := GetRectangleAndSquareArea(fmt.Sprintf("localhost:%d", transport.Port), request)

		assert.NoError(t, err)
		assert.Equal(t, uint32(0), area.GetCode())

		return err
	}

	pact.RunTest(t, *c, *cp)

}
