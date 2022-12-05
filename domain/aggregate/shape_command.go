package aggregate

import (
	"example2/infra"
	"github.com/google/uuid"
)

type newShapeCommand struct {
	nature     string
	dimensions []float32
}

func (n *newShapeCommand) Apply(apply ApplyShapeCommand) (Shape, []infra.Event) {
	return apply.ApplyNewShapeCommand(*n)
}

func newCreationShapeCommand(nature string, dimensions []float32) (*newShapeCommand, error) {
	command := new(newShapeCommand)
	command.nature = nature
	command.dimensions = dimensions
	return command, nil
}

type newStretchCommand struct {
	id        uuid.UUID
	stretchBy float32
}

func (n *newStretchCommand) Apply(apply ApplyShapeCommand) (Shape, []infra.Event) {
	return apply.ApplyNewStretchCommand(*n)
}

func newStrechShapeCommand(id uuid.UUID, stretchBy float32) *newStretchCommand {
	command := new(newStretchCommand)
	command.id = id
	command.stretchBy = stretchBy
	return command
}
