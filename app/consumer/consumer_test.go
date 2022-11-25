package consumer

import (
	ac "example2/app/proto"
	"example2/app/utils"
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

var dir, _ = filepath.Abs("../proto/app_area_calculator.proto")

func TestCalculateClient(t *testing.T) {
	grpcInteraction := `{
		"pact:proto": "` + dir + `",
		"pact:proto-service": "Calculator/calculateMultiV2",
		"pact:content-type": "application/protobuf",
		"request": {
			"shapes": [
				{
					"shape": "rectangle",
					"dimensions": [ "matching(number, 3)", "matching(number, 4)"]
				}
			]
		},
		"response": {
			"value": [ "matching(number, 12)"]
		}
	}`

	var c = new(utils.ContractTest)
	c.GrpcInteraction = grpcInteraction
	c.Description = "calculate rectangle area request"
	c.F = func(transport message.TransportConfig, m message.SynchronousMessage) error {
		dimensions := []float32{3, 4}
		shape := "rectangle"
		rectangle := ac.ShapeMessageV2{Shape: shape, Dimensions: dimensions}
		request := &ac.AreaRequestV2{Shapes: []*ac.ShapeMessageV2{&rectangle}}
		area, err := GetRectangleAndSquareArea(fmt.Sprintf("localhost:%d", transport.Port), request)

		assert.NoError(t, err)
		assert.Equal(t, float32(12.0), area[0])

		return err
	}

	utils.RunTest(t, *c, *cp)

}

func TestCalculateClientNoArea(t *testing.T) {
	grpcInteraction := `{
		"pact:proto": "` + dir + `",
		"pact:proto-service": "Calculator/calculateMultiV2",
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
		request := &ac.AreaRequestV2{Shapes: []*ac.ShapeMessageV2{}}
		area, err := GetRectangleAndSquareArea(fmt.Sprintf("localhost:%d", transport.Port), request)
		fmt.Println(area)
		assert.NoError(t, err)
		assert.Equal(t, float32(0), area[0])

		return err
	}

	utils.RunTest(t, *c, *cp)

}
