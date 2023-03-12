package shape

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/shape-app/eventstore"
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
	provider     eventstore.Provider
	eventFactory *factoryEvents
}

func newApplyShapeCommand(provider eventstore.Provider) ApplyShapeCommand {
	a := new(ApplyShapeCommandImpl)
	a.provider = provider
	a.eventFactory = newEventFactory()
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
	shape := a.createShape(events[0])
	for _, e := range events {
		domainEvent := a.eventFactory.newDeserializedEvent(uuid, e)
		shape = domainEvent.(Event).ApplyOn(shape)
	}
	return shape, nil
}

func (a ApplyShapeCommandImpl) createShape(createdEvent eventstore.DomainEvent) Shape {
	a.checkEventName(createdEvent.Name())
	initialEvent := a.eventFactory.newDeserializedEvent(createdEvent.AggregateId(), createdEvent).(*Created)
	shape, _ := newShapeBuilder().withNature(initialEvent.Nature).withId(initialEvent.AggregateId())
	return shape
}

func (a ApplyShapeCommandImpl) checkEventName(actualEventName string) {
	notOk := actualEventName != "SHAPE_CREATED"
	if notOk {
		panic(fmt.Errorf("unexpected event: %s", actualEventName))
	}
}
