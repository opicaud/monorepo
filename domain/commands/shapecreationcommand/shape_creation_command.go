package shapecreationcommand

import (
	"errors"
	"example2/domain/commands"
	"example2/domain/valueobject"
)

type fullShapeCommand struct {
	shape valueobject.Shape
}

func NewFullShapeCommand(shape valueobject.Shape) (commands.Command, error) {
	if nil == shape {
		return createCommand(nil), errors.New("shape is mandatory")
	}
	return createCommand(shape), nil
}

func createCommand(shape valueobject.Shape) fullShapeCommand {
	command := new(fullShapeCommand)
	command.shape = shape
	return *command
}

func (f fullShapeCommand) Execute() error {
	return nil
}
