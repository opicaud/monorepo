package aggregate

import (
	"example2/domain/commands"
)

type newShapeCommand struct {
	nature     string
	dimensions []float32
}

func newCreationShapeCommand(nature string, dimensions []float32) (commands.Command, error) {
	return createCommand(nature, dimensions), nil
}

func createCommand(nature string, dimensions []float32) newShapeCommand {
	command := new(newShapeCommand)
	command.nature = nature
	command.dimensions = dimensions
	return *command
}
