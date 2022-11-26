package commands

import "errors"

type fullShapeCommand struct {
	nature     string
	dimensions []float32
}

func newFullShapeCommand(n string, d ...float32) (*fullShapeCommand, error) {
	if nil == d {
		return nil, errors.New("dimensions are mandatory")
	}
	return createCommand(n, d), nil
}

func createCommand(n string, d []float32) *fullShapeCommand {
	command := new(fullShapeCommand)
	command.nature = n
	command.dimensions = d
	return command
}

func (f fullShapeCommand) Execute() error {
	return nil
}
