package commands

import (
	"example2/domain/aggregate"
	"example2/domain/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullShapeCommand(t *testing.T) {
	command, err := newFullShapeCommand("nature", 1, 2)
	assert.NoError(t, err)
	assert.Equal(t, "nature", command.nature)
	assert.Equal(t, []float32{1, 2}, command.dimensions)
	assert.IsType(t, &aggregate.ShapeBuilder{}, command.builder)
}

func TestFullShapeCommandErrorWhenNoDimensionsProvided(t *testing.T) {
	_, err := newFullShapeCommand("nature")
	assert.Error(t, err)
}

func TestExecuteFullShapeCommand(t *testing.T) {
	builderForTest := utils.FakeShapeBuilder{}
	command := createCommandWithCustomBuilder("nature", []float32{2, 3}, &builderForTest)
	command.Execute()
	builderForTest.Mock.AssertCalled(t, "Area")
}
