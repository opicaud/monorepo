package shape

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/opicaud/monorepo/shape-app/cqrs"
)

type CreationCommand struct {
	nature     string
	dimensions []float32
}

func (n *CreationCommand) Execute(apply CommandApplier) ([]pkg.DomainEvent, error) {
	return apply.ApplyCreationCommand(*n)
}

func newCreationShapeCommand(nature string, dimensions []float32) *CreationCommand {
	command := new(CreationCommand)
	command.nature = nature
	command.dimensions = dimensions
	return command
}

type StretchCommand struct {
	id        uuid.UUID
	stretchBy float32
}

func (n *StretchCommand) Execute(apply CommandApplier) ([]pkg.DomainEvent, error) {
	return apply.ApplyStretchCommand(*n)
}

func newStretchShapeCommand(id uuid.UUID, stretchBy float32) *StretchCommand {
	command := new(StretchCommand)
	command.id = id
	command.stretchBy = stretchBy
	return command
}

type commandApplier struct {
	eventsFramework pkg.Provider
}

func NewShapeCommandApplier(eventsFramework pkg.Provider) CommandApplier {
	a := new(commandApplier)
	a.eventsFramework = eventsFramework
	return a
}

func (commandApplier) ApplyCreationCommand(command CreationCommand) ([]pkg.DomainEvent, error) {
	shape, err := newShapeBuilder().withNature(command.nature).withId(uuid.New())
	if err != nil {
		return nil, err
	}
	return []pkg.DomainEvent{shape.HandleCreationCommand(command)}, nil
}

func (a commandApplier) ApplyStretchCommand(command StretchCommand) ([]pkg.DomainEvent, error) {
	shape, err := a.loadShapeFromEventStore(command.id)
	if err != nil {
		return nil, err
	}

	return []pkg.DomainEvent{shape.HandleStretchCommand(command)}, nil

}

func (a commandApplier) loadShapeFromEventStore(uuid uuid.UUID) (Shape, error) {
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

func (a commandApplier) createShape(shapeEventFactory shapeEventFactory, createdEvent pkg.DomainEvent) Shape {
	a.checkEventName(createdEvent.Name())

	initialEvent := shapeEventFactory.newDeserializedEvent(createdEvent.AggregateId(), createdEvent).(*Created)
	shape, _ := newShapeBuilder().withNature(initialEvent.Nature).withId(initialEvent.AggregateId())
	return shape
}

func (a commandApplier) checkEventName(actualEventName string) {
	notOk := actualEventName != "SHAPE_CREATED"
	if notOk {
		panic(fmt.Errorf("unexpected event: %s", actualEventName))
	}
}

func NewCommandHandlerBuilder() *cqrs.CommandHandlerBuilder[CommandApplier] {
	return &cqrs.CommandHandlerBuilder[CommandApplier]{}
}

type Event interface {
	ApplyOn(shape Shape) Shape
}
