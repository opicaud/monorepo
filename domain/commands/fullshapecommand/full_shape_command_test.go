package fullshapecommand

import (
	"example2/domain/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullShapeCommand(t *testing.T) {
	shape := utils.MockShape{}
	command, err := NewFullShapeCommand(&shape)
	assert.NoError(t, err)
	assert.Equal(t, &shape, command.shape)
}

func TestFullShapeCommandErrorWhenNoShapeProvided(t *testing.T) {
	_, err := NewFullShapeCommand(nil)
	assert.Error(t, err)
}

func TestExecuteFullShapeCommand(t *testing.T) {
	shape := utils.CreateAMockShape()
	command, _ := NewFullShapeCommand(&shape)
	command.Execute()
	shape.Mock.AssertCalled(t, "Area")
}
