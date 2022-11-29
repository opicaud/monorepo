package utils

import (
	"example2/domain/commands/factory"
	"example2/domain/commands/fullshapecommand"
	"example2/domain/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThemAll(t *testing.T) {
	handler := fullshapecommand.NewFullShapeCommandHandler(utils.NewFakeRepository())
	factory := factory.NewFactory()
	command, err := factory.CreateAFullShapeCommand("rectangle", 1, 2)
	assert.NoError(t, err)
	handler.Execute(command)
}
