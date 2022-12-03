package valueobject

import (
	"example2/domain/commands"
)

func (f *factoryOfShapeCommand) NewCreationShapeCommand(nature string, dimensions ...float32) (commands.Command, error) {
	shape, err := f.builder.CreateAShape(nature).WithDimensions(dimensions)
	command, _ := NewCreationShapeCommand(shape)
	return command, err
}

func (f *factoryOfShapeCommand) NewCreationShapeCommand2(nature string, dimensions ...float32) (commands.Command, error) {
	command, _ := NewCreationShapeCommand2(nature, dimensions)
	return command, nil
}

type factoryOfShapeCommand struct {
	builder IShapeBuilder
}

func newFactoryWithCustomBuilder(builder IShapeBuilder) *factoryOfShapeCommand {
	factoryOfShapeCommand := new(factoryOfShapeCommand)
	factoryOfShapeCommand.builder = builder
	return factoryOfShapeCommand
}

func NewFactory() *factoryOfShapeCommand {
	factoryOfShapeCommand := new(factoryOfShapeCommand)
	factoryOfShapeCommand.builder = NewShapeBuilder()
	return factoryOfShapeCommand
}

type IFactoryOfShapeCommand interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) (commands.Command, error)
}
