package aggregate

import (
	"example2/domain/commands"
	"github.com/google/uuid"
)

type newShapeCommand struct {
	nature     string
	dimensions []float32
}

func newCreationShapeCommand(nature string, dimensions []float32) (commands.Command, error) {
	command := new(newShapeCommand)
	command.nature = nature
	command.dimensions = dimensions
	return *command, nil
}

type newStretchCommand struct {
	id        uuid.UUID
	stretchBy float32
}

func newStrechShapeCommand(id uuid.UUID, stretchBy float32) commands.Command {
	command := new(newStretchCommand)
	command.id = id
	command.stretchBy = stretchBy
	return *command
}
