package commands

func (f *factoryOfShapeCommand) CreateAFullShapeCommand(nature string, dimensions ...float32) (Command, error) {
	command, err := newFullShapeCommand(nature, dimensions...)
	return command, err
}

type factoryOfShapeCommand struct{}

func New() *factoryOfShapeCommand {
	return new(factoryOfShapeCommand)
}

type IFactoryOfShapeCommand interface {
	CreateAFullShapeCommand(nature string, dimensions ...float32) (Command, error)
}
