package factories

import "example2/domain/commands"

func (f *factoryOfShapeCommand) CreateAFullShapeCommand(nature string, dimensions ...float32) (commands.Command, error) {
	command, err := commands.NewFullShapeCommand(nature, dimensions...)
	return command, err
}

type factoryOfShapeCommand struct{}

func New() *factoryOfShapeCommand {
	return new(factoryOfShapeCommand)
}

type IFactoryOfShapeCommand interface {
	CreateAFullShapeCommand(nature string, dimensions ...float32) (commands.Command, error)
}
