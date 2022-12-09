package aggregate

import (
	"example2/infra"
	"github.com/google/uuid"
	"github.com/smartystreets/assertions"
)

type newShapeCommand struct {
	nature     string
	dimensions []float32
}

func (n *newShapeCommand) Apply(apply ApplyShapeCommand) []infra.Event {
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

func (n *newStretchCommand) Apply(apply ApplyShapeCommand) []infra.Event {
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

func (ApplyShapeCommandImpl) ApplyNewShapeCommand(command newShapeCommand) []infra.Event {
	shape, err := newShapeBuilder().createAShape(command.nature).withId(uuid.New())
	if err != nil {
		panic(err)
	}
	shapeCreatedEvent, areaShapeCalculated := shape.HandleNewShape(command)
	return []infra.Event{shapeCreatedEvent, areaShapeCalculated}
}

func (a ApplyShapeCommandImpl) ApplyNewStretchCommand(command newStretchCommand) []infra.Event {
	events := a.provider.Load(command.id)

	assertions.ShouldImplement(events[0], ShapeCreatedEvent{})
	initialEvent := events[0].(ShapeCreatedEvent)

	shape, _ := newShapeBuilder().createAShape(initialEvent.Nature).withId(initialEvent.id)

	for _, e := range events {
		shape = e.(ShapeEvent).Apply(shape)
	}

	areaShapeCalculated := shape.HandleStretchCommand(command)
	return []infra.Event{areaShapeCalculated}

}
