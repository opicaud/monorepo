package shape

import (
	"github.com/google/uuid"
)

func (f *FactoryOfShapeCommand) NewCreationShapeCommand(nature string, dimensions ...float32) Command[CommandApplier] {
	command := newCreationShapeCommand(nature, dimensions)
	return command
}

func (f *FactoryOfShapeCommand) NewStretchShapeCommand(id uuid.UUID, stretchBy float32) Command[CommandApplier] {
	command := newStretchShapeCommand(id, stretchBy)
	return command
}

func NewFactory() *FactoryOfShapeCommand {
	return new(FactoryOfShapeCommand)
}

type FactoryOfShapeCommand struct{}

type IFactoryOfShapeCommand interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) (Command[CommandApplier], error)
	NewStretchShapeCommand(id uuid.UUID, stretchBy float32) Command[CommandApplier]
}
