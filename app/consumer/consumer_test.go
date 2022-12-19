package consumer

import (
	"fmt"
	"path/filepath"
	"testing"
	ac "trackclear.be/example/shape/app/proto"
	"trackclear.be/example/shape/app/utils"

	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"
)

var cp = &utils.ConsumerAndProvider{
	Consumer: "grpc-consumer-go",
	Provider: "area-calculator-provider",
}

var dir, _ = filepath.Abs("../proto/app_shape.proto")

func TestCreateShape(t *testing.T) {
	grpcInteraction := `{
		"pact:proto": "` + dir + `",
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

	var c = new(utils.ContractTest)
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

	utils.RunTest(t, *c, *cp)

}
