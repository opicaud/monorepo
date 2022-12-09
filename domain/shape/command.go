package shape

import (
	"example2/infra"
	"github.com/google/uuid"
	"github.com/smartystreets/assertions"
)

type newShapeCommand struct {
	nature     string
	dimensions []float32
}

func (n *newShapeCommand) Apply(apply ApplyShapeCommand) ([]infra.Event, error) {
	return apply.ApplyNewShapeCommand(*n)
}

func newCreationShapeCommand(nature string, dimensions []float32) *newShapeCommand {
	command := new(newShapeCommand)
	command.nature = nature
	command.dimensions = dimensions
	return command
}

type newStretchCommand struct {
	id        uuid.UUID
	stretchBy float32
}

func (n *newStretchCommand) Apply(apply ApplyShapeCommand) ([]infra.Event, error) {
	return apply.ApplyNewStretchCommand(*n)
}

func newStrechShapeCommand(id uuid.UUID, stretchBy float32) *newStretchCommand {
	command := new(newStretchCommand)
	command.id = id
	command.stretchBy = stretchBy
	return command
}

type ApplyShapeCommandImpl struct {
	provider infra.EventStore
}

func newApplyShapeCommand(provider infra.Provider) ApplyShapeCommand {
	a := new(ApplyShapeCommandImpl)
	a.provider = &provider
	return a
}

func (ApplyShapeCommandImpl) ApplyNewShapeCommand(command newShapeCommand) ([]infra.Event, error) {
	shape, err := newShapeBuilder().createAShape(command.nature).withId(uuid.New())
	if err != nil {
		return nil, err
	}
	return []infra.Event{shape.HandleNewShape(command)}, nil
}

func (a ApplyShapeCommandImpl) ApplyNewStretchCommand(command newStretchCommand) ([]infra.Event, error) {
	shape, err := a.loadShapeFromEventStore(command.id)
	if err != nil {
		return nil, err
	}

	return []infra.Event{shape.HandleStretchCommand(command)}, nil

}

func (a ApplyShapeCommandImpl) loadShapeFromEventStore(uuid uuid.UUID) (Shape, error) {
	events, err := a.provider.Load(uuid)
	if err != nil {
		return nil, err
	}
	assertions.ShouldImplement(events[0], ShapeCreated{})
	initialEvent := events[0].(ShapeCreated)
	shape, _ := newShapeBuilder().createAShape(initialEvent.Nature).withId(initialEvent.id)
	for _, e := range events {
		shape = e.(ShapeEvent).Apply(shape)
	}
	return shape, nil
}
