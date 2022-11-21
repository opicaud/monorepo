package consumer

import (
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
		area, err := GetRectangleAndSquareArea(fmt.Sprintf("localhost:%d", transport.Port))

		assert.NoError(t, err)
		assert.Equal(t, float32(12.0), area[0])
		assert.Equal(t, float32(9.0), area[1])

		return err
	}

	utils.RunTest(t, *c, *cp)

}
