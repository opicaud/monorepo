package aggregate

import (
	"example2/domain/commands"
	"github.com/google/uuid"
)

func (f *factoryOfShapeCommand) NewCreationShapeCommand(nature string, dimensions ...float32) (commands.Command, error) {
	command, _ := newCreationShapeCommand(nature, dimensions)
	return command, nil
}

func (f *factoryOfShapeCommand) NewStretchShapeCommand(id uuid.UUID, stretchBy float32) commands.Command {
	command := newStrechShapeCommand(id, stretchBy)
	return command
}

func NewFactory() factoryOfShapeCommand {
	return factoryOfShapeCommand{}
}

type factoryOfShapeCommand struct{}

type IFactoryOfShapeCommand interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) (commands.Command, error)
	NewStretchShapeCommand(id uuid.UUID, stretchBy float32) commands.Command
}
