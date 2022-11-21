package consumer

import (
	ac "example2/proto"
	"example2/utils"
	"fmt"
	"path/filepath"
	"testing"

	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"
)

var cp = &utils.ConsumerAndProvider{
	Consumer: "grpc-consumer-go",
	Provider: "area-calculator-provider",
}

var dir, _ = filepath.Abs("./proto/area_calculator.proto")

func TestCalculateClient(t *testing.T) {
	grpcInteraction := `{
		"pact:proto": "` + dir + `",
		"pact:proto-service": "Calculator/calculateMulti",
		"pact:content-type": "application/protobuf",
		"request": {
			"shapes": [
				{
					"rectangle": {
						"length": "matching(number, 3)",
						"width": "matching(number, 4)"
					}
				},
				{
					"square": {
						"edge_length": "matching(number, 3)"
					}
				}
			]
		},
		"response": {
			"value": [ "matching(number, 12)", "matching(number, 9)" ]
		}
	}`

	var c = new(utils.ContractTest)
	c.GrpcInteraction = grpcInteraction
	c.Description = "calculate rectangle area request"
	c.F = func(transport message.TransportConfig, m message.SynchronousMessage) error {
		rectangle := ac.ShapeMessage{Shape: &ac.ShapeMessage_Rectangle{Rectangle: &ac.Rectangle{Length: 3, Width: 4}}}
		square := ac.ShapeMessage{Shape: &ac.ShapeMessage_Square{Square: &ac.Square{EdgeLength: 3}}}
		request := &ac.AreaRequest{Shapes: []*ac.ShapeMessage{&rectangle, &square}}
		area, err := GetRectangleAndSquareArea(fmt.Sprintf("localhost:%d", transport.Port), *request)

		assert.NoError(t, err)
		assert.Equal(t, float32(12.0), area[0])
		assert.Equal(t, float32(9.0), area[1])

		return err
	}

	utils.RunTest(t, *c, *cp)

}

func TestCalculateClientNoArea(t *testing.T) {
	grpcInteraction := `{
		"pact:proto": "` + dir + `",
		"pact:proto-service": "Calculator/calculateMulti",
		"pact:content-type": "application/protobuf",
		"request": {
			"shapes": []
		},
		"response": {
			"value": [ "matching(number, 0)" ]
		}
	}`

	var c = new(utils.ContractTest)
	c.GrpcInteraction = grpcInteraction
	c.Description = "calculate no area"
	c.F = func(transport message.TransportConfig, m message.SynchronousMessage) error {
		request := &ac.AreaRequest{Shapes: []*ac.ShapeMessage{}}
		area, err := GetRectangleAndSquareArea(fmt.Sprintf("localhost:%d", transport.Port), *request)

		assert.NoError(t, err)
		assert.Equal(t, float32(0), area[0])

		return err
	}

	utils.RunTest(t, *c, *cp)

}
