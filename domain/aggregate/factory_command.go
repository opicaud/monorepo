package aggregate

import (
	"github.com/google/uuid"
)

func (f *factoryOfShapeCommand) NewCreationShapeCommand(nature string, dimensions ...float32) ShapeCommand {
	command := newCreationShapeCommand(nature, dimensions)
	return command
}

func (f *factoryOfShapeCommand) NewStretchShapeCommand(id uuid.UUID, stretchBy float32) ShapeCommand {
	command := newStrechShapeCommand(id, stretchBy)
	return command
}

func NewFactory() factoryOfShapeCommand {
	return factoryOfShapeCommand{}
}

type factoryOfShapeCommand struct{}

type IFactoryOfShapeCommand interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) (ShapeCommand, error)
	NewStretchShapeCommand(id uuid.UUID, stretchBy float32) ShapeCommand
}
