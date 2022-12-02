package shapecreationcommand

import (
	"errors"
	"example2/domain/commands"
	"example2/domain/valueobject"
)

type newShapeCommand struct {
	shape valueobject.Shape
}

func NewCreationShapeCommand(shape valueobject.Shape) (commands.Command, error) {
	if nil == shape {
		return createCommand(nil), errors.New("shape is mandatory")
	}
	return createCommand(shape), nil
}

func createCommand(shape valueobject.Shape) newShapeCommand {
	command := new(newShapeCommand)
	command.shape = shape
	return *command
}

func (f newShapeCommand) Execute() error {
	return nil
}
