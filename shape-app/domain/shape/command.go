package shape

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
)

type newShapeCommand struct {
	nature     string
	dimensions []float32
}

func (n *newShapeCommand) Execute(apply ShapeCommandApplier) ([]pkg.DomainEvent, error) {
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

func (n *newStretchCommand) Execute(apply ShapeCommandApplier) ([]pkg.DomainEvent, error) {
	return apply.ApplyNewStretchCommand(*n)
}

func newStretchShapeCommand(id uuid.UUID, stretchBy float32) *newStretchCommand {
	command := new(newStretchCommand)
	command.id = id
	command.stretchBy = stretchBy
	return command
}

type ShapeCommandApplierImpl struct {
	eventsFramework pkg.Provider
}

func NewShapeCommandApplier(eventsFramework pkg.Provider) ShapeCommandApplier {
	a := new(ShapeCommandApplierImpl)
	a.eventsFramework = eventsFramework
	return a
}

func (ShapeCommandApplierImpl) ApplyNewShapeCommand(command newShapeCommand) ([]pkg.DomainEvent, error) {
	shape, err := newShapeBuilder().withNature(command.nature).withId(uuid.New())
	if err != nil {
		return nil, err
	}
	return []pkg.DomainEvent{shape.HandleNewShape(command)}, nil
}

func (a ShapeCommandApplierImpl) ApplyNewStretchCommand(command newStretchCommand) ([]pkg.DomainEvent, error) {
	shape, err := a.loadShapeFromEventStore(command.id)
	if err != nil {
		return nil, err
	}

	return []pkg.DomainEvent{shape.HandleStretchCommand(command)}, nil

}

func (a ShapeCommandApplierImpl) loadShapeFromEventStore(uuid uuid.UUID) (Shape, error) {
	events, err := a.eventsFramework.Load(uuid)
	if err != nil {
		return nil, err
	}
	shapeEventsFactory := newShapeEventFactory()
	shape := a.createShape(shapeEventsFactory, events[0])
	for _, e := range events {
		shapeEvent := shapeEventsFactory.newDeserializedEvent(uuid, e)
		shape = shapeEvent.(Event).ApplyOn(shape)
	}
	return shape, nil
}

func (a ShapeCommandApplierImpl) createShape(shapeEventFactory shapeEventFactory, createdEvent pkg.DomainEvent) Shape {
	a.checkEventName(createdEvent.Name())

	initialEvent := shapeEventFactory.newDeserializedEvent(createdEvent.AggregateId(), createdEvent).(*Created)
	shape, _ := newShapeBuilder().withNature(initialEvent.Nature).withId(initialEvent.AggregateId())
	return shape
}

func (a ShapeCommandApplierImpl) checkEventName(actualEventName string) {
	notOk := actualEventName != "SHAPE_CREATED"
	if notOk {
		panic(fmt.Errorf("unexpected event: %s", actualEventName))
	}
}
