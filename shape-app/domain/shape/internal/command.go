package internal

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
)

type CreationCommand struct {
	nature     string
	dimensions []float32
}

func (n *CreationCommand) Execute(apply CommandApplier) ([]pkg.DomainEvent, error) {
	return apply.ApplyCreationCommand(*n)
}

func NewCreationShapeCommand(nature string, dimensions []float32) *CreationCommand {
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

func NewStretchShapeCommand(id uuid.UUID, stretchBy float32) *StretchCommand {
	command := new(StretchCommand)
	command.id = id
	command.stretchBy = stretchBy
	return command
}

type StandardCommandApplier struct {
	eventsFramework pkg.Provider
}

func NewShapeCommandApplier(eventsFramework pkg.Provider) CommandApplier {
	a := new(StandardCommandApplier)
	a.eventsFramework = eventsFramework
	return a
}

func (StandardCommandApplier) ApplyCreationCommand(command CreationCommand) ([]pkg.DomainEvent, error) {
	shape, err := newShapeBuilder().withNature(command.nature).withId(uuid.New())
	if err != nil {
		return nil, err
	}
	return []pkg.DomainEvent{shape.HandleCreationCommand(command)}, nil
}

func (a StandardCommandApplier) ApplyStretchCommand(command StretchCommand) ([]pkg.DomainEvent, error) {
	shape, err := a.loadShapeFromEventStore(command.id)
	if err != nil {
		return nil, err
	}

	return []pkg.DomainEvent{shape.HandleStretchCommand(command)}, nil

}

func (a StandardCommandApplier) loadShapeFromEventStore(uuid uuid.UUID) (Shape, error) {
	events, err := a.eventsFramework.Load(uuid)
	if err != nil {
		return nil, err
	}
	shapeEventsFactory := NewShapeEventFactory()
	shape := a.createShape(events[0])
	for _, e := range events {
		shapeEvent := shapeEventsFactory.NewDeserializedEvent(uuid, e)
		shape = shapeEvent.(Event).ApplyOn(shape)
	}
	return shape, nil
}

func (a StandardCommandApplier) createShape(createdEvent pkg.DomainEvent) Shape {
	a.checkEventName(createdEvent.Name())
	n := nature{}
	_ = json.Unmarshal(createdEvent.Data(), &n)
	shape, _ := newShapeBuilder().withNature(n.Nature).withId(createdEvent.AggregateId())
	return shape
}

type nature struct {
	Nature string
}

func (a StandardCommandApplier) checkEventName(actualEventName string) {
	notOk := actualEventName != "SHAPE_CREATED"
	if notOk {
		panic(fmt.Errorf("unexpected event: %s", actualEventName))
	}
}

type Event interface {
	ApplyOn(shape Shape) Shape
}
