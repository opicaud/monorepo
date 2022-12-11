package shape

import (
	"github.com/google/uuid"
)

func (f *FactoryOfShapeCommand) NewCreationShapeCommand(nature string, dimensions ...float32) Command {
	command := newCreationShapeCommand(nature, dimensions)
	return command
}

func (f *FactoryOfShapeCommand) NewStretchShapeCommand(id uuid.UUID, stretchBy float32) Command {
	command := newStretchShapeCommand(id, stretchBy)
	return command
}

func NewFactory() *FactoryOfShapeCommand {
	return new(FactoryOfShapeCommand)
}

type FactoryOfShapeCommand struct{}

type IFactoryOfShapeCommand interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) (Command, error)
	NewStretchShapeCommand(id uuid.UUID, stretchBy float32) Command
}
