package commands

import (
	"errors"
	"example2/domain/valueobject"
)

type fullShapeCommand struct {
	shape valueobject.Shape
}

func newFullShapeCommand(shape valueobject.Shape) (*fullShapeCommand, error) {
	if nil == shape {
		return nil, errors.New("shape is mandatory")
	}
	return createCommand(shape), nil
}

func createCommand(shape valueobject.Shape) *fullShapeCommand {
	command := new(fullShapeCommand)
	command.shape = shape
	return command
}

func (f fullShapeCommand) Execute() error {
	err := f.shape.Area()
	return err
}
