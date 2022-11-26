package commands

import "example2/domain/valueobject"

func (f *factoryOfShapeCommand) CreateAFullShapeCommand(nature string, dimensions ...float32) (Command, error) {
	shape, err := f.builder.CreateAShape(nature).WithDimensions(dimensions)
	command, _ := newFullShapeCommand(shape)
	return command, err
}

type factoryOfShapeCommand struct {
	builder valueobject.IShapeBuilder
}

func New(builder valueobject.IShapeBuilder) *factoryOfShapeCommand {
	factoryOfShapeCommand := new(factoryOfShapeCommand)
	factoryOfShapeCommand.builder = builder
	return factoryOfShapeCommand
}

type IFactoryOfShapeCommand interface {
	CreateAFullShapeCommand(nature string, dimensions ...float32) (Command, error)
}
