package valueobject

import (
	"example2/domain/commands"
)

func (f *factoryOfShapeCommand) NewCreationShapeCommand(nature string, dimensions ...float32) (commands.Command, error) {
	command, _ := newCreationShapeCommand(nature, dimensions)
	return command, nil
}

func NewFactory() factoryOfShapeCommand {
	return factoryOfShapeCommand{}
}

type factoryOfShapeCommand struct{}

type IFactoryOfShapeCommand interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) (commands.Command, error)
}
