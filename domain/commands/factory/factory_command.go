package factory

import (
	"example2/domain/commands"
	"example2/domain/commands/fullshapecommand"
	"example2/domain/valueobject"
)

func (f *factoryOfShapeCommand) NewShapeCreationCommand(nature string, dimensions ...float32) (commands.Command, error) {
	shape, err := f.builder.CreateAShape(nature).WithDimensions(dimensions)
	command, _ := fullshapecommand.NewFullShapeCommand(shape)
	return command, err
}

type factoryOfShapeCommand struct {
	builder valueobject.IShapeBuilder
}

func NewFactoryWithCustomBuilder(builder valueobject.IShapeBuilder) *factoryOfShapeCommand {
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
	NewShapeCreationCommand(nature string, dimensions ...float32) (commands.Command, error)
}
