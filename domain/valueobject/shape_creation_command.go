package valueobject

import (
	"errors"
	"example2/domain/commands"
)

type newShapeCommand struct {
	shape      Shape
	nature     string
	dimensions []float32
}

func NewCreationShapeCommand(shape Shape) (commands.Command, error) {
	if nil == shape {
		return createCommand(nil), errors.New("shape is mandatory")
	}
	return createCommand(shape), nil
}

func NewCreationShapeCommand2(nature string, dimensions []float32) (commands.Command, error) {
	return createCommand2(nature, dimensions), nil
}

func createCommand(shape Shape) newShapeCommand {
	command := new(newShapeCommand)
	command.shape = shape
	return *command
}

func createCommand2(nature string, dimensions []float32) newShapeCommand {
	command := new(newShapeCommand)
	command.nature = nature
	command.dimensions = dimensions
	return *command
}

func (f newShapeCommand) Execute() error {
	return nil
}
