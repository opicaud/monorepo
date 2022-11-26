package commands

import (
	"errors"
	"example2/domain/aggregate"
)

type fullShapeCommand struct {
	nature     string
	dimensions []float32
	builder    aggregate.IShapeBuilder
}

func newFullShapeCommand(n string, d ...float32) (*fullShapeCommand, error) {
	if nil == d {
		return nil, errors.New("dimensions are mandatory")
	}
	return createCommand(n, d), nil
}

func createCommandWithCustomBuilder(n string, d []float32, builder aggregate.IShapeBuilder) *fullShapeCommand {
	command := createCommand(n, d)
	command.builder = builder
	return command
}

func createCommand(n string, d []float32) *fullShapeCommand {
	command := new(fullShapeCommand)
	command.nature = n
	command.dimensions = d
	command.builder = aggregate.NewShapeBuilder()
	return command
}

func (f fullShapeCommand) Execute() error {
	shape, err := f.builder.CreateAShape(f.nature).WithDimensions(f.dimensions)
	shape.Area()
	return err
}
