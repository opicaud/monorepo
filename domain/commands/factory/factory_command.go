package factory

import (
	"example2/domain/commands"
	"example2/domain/commands/shapecreationcommand"
	"example2/domain/valueobject"
)

func (f *factoryOfShapeCommand) NewCreationShapeCommand(nature string, dimensions ...float32) (commands.Command, error) {
	shape, err := f.builder.CreateAShape(nature).WithDimensions(dimensions)
	command, _ := shapecreationcommand.NewCreationShapeCommand(shape)
	return command, err
}

type factoryOfShapeCommand struct {
	builder valueobject.IShapeBuilder
}

func newFactoryWithCustomBuilder(builder valueobject.IShapeBuilder) *factoryOfShapeCommand {
	factoryOfShapeCommand := new(factoryOfShapeCommand)
	factoryOfShapeCommand.builder = builder
	return factoryOfShapeCommand
}

func NewFactory() *factoryOfShapeCommand {
	factoryOfShapeCommand := new(factoryOfShapeCommand)
	factoryOfShapeCommand.builder = valueobject.NewShapeBuilder()
	return factoryOfShapeCommand
}

type IFactoryOfShapeCommand interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) (commands.Command, error)
}
