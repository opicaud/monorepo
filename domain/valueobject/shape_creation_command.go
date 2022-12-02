package valueobject

import (
	"errors"
	"example2/domain/commands"
)

type newShapeCommand struct {
	shape Shape
}

func NewCreationShapeCommand(shape Shape) (commands.Command, error) {
	if nil == shape {
		return createCommand(nil), errors.New("shape is mandatory")
	}
	return createCommand(shape), nil
}

func createCommand(shape Shape) newShapeCommand {
	command := new(newShapeCommand)
	command.shape = shape
	return *command
}

func (f newShapeCommand) Execute() error {
	return nil
}
