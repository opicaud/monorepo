package shape

import (
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/shape-app/eventstore"
	"github.com/smartystreets/assertions"
)

type newShapeCommand struct {
	nature     string
	dimensions []float32
}

func (n *newShapeCommand) Execute(apply ApplyShapeCommand) ([]eventstore.DomainEvent, error) {
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

func (n *newStretchCommand) Execute(apply ApplyShapeCommand) ([]eventstore.DomainEvent, error) {
	return apply.ApplyNewStretchCommand(*n)
}

func newStretchShapeCommand(id uuid.UUID, stretchBy float32) *newStretchCommand {
	command := new(newStretchCommand)
	command.id = id
	command.stretchBy = stretchBy
	return command
}

type ApplyShapeCommandImpl struct {
	provider eventstore.Provider
}

func newApplyShapeCommand(provider eventstore.Provider) ApplyShapeCommand {
	a := new(ApplyShapeCommandImpl)
	a.provider = provider
	return a
}

func (ApplyShapeCommandImpl) ApplyNewShapeCommand(command newShapeCommand) ([]eventstore.DomainEvent, error) {
	shape, err := newShapeBuilder().withNature(command.nature).withId(uuid.New())
	if err != nil {
		return nil, err
	}
	return []eventstore.DomainEvent{shape.HandleNewShape(command)}, nil
}

func (a ApplyShapeCommandImpl) ApplyNewStretchCommand(command newStretchCommand) ([]eventstore.DomainEvent, error) {
	shape, err := a.loadShapeFromEventStore(command.id)
	if err != nil {
		return nil, err
	}

	return []eventstore.DomainEvent{shape.HandleStretchCommand(command)}, nil

}

func (a ApplyShapeCommandImpl) loadShapeFromEventStore(uuid uuid.UUID) (Shape, error) {
	events, err := a.provider.Load(uuid)
	if err != nil {
		return nil, err
	}
	assertions.ShouldImplement(events[0], Created{})
	initialEvent := events[0].(Created)
	shape, _ := newShapeBuilder().withNature(initialEvent.Nature).withId(initialEvent.id)
	for _, e := range events {
		shape = e.(Event).ApplyOn(shape)
	}
	return shape, nil
}