package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	cqrs "github.com/opicaud/monorepo/cqrs/pkg/v3beta1"
)

type CreationCommand struct {
	nature     string
	dimensions []float32
}

func (n *CreationCommand) Execute(apply CommandApplier) ([]cqrs.DomainEvent, error) {
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

func (n *StretchCommand) Execute(apply CommandApplier) ([]cqrs.DomainEvent, error) {
	return apply.ApplyStretchCommand(*n)
}

func NewStretchShapeCommand(id uuid.UUID, stretchBy float32) *StretchCommand {
	command := new(StretchCommand)
	command.id = id
	command.stretchBy = stretchBy
	return command
}

type StandardCommandApplier struct {
	eventStore cqrs.EventStore
}

func NewShapeCommandApplier(eventStore cqrs.EventStore) CommandApplier {
	a := new(StandardCommandApplier)
	a.eventStore = eventStore
	return a
}

func (StandardCommandApplier) ApplyCreationCommand(command CreationCommand) ([]cqrs.DomainEvent, error) {
	shape, err := newShapeBuilder().withNature(command.nature).withId(uuid.New())
	if err != nil {
		return nil, err
	}
	return []cqrs.DomainEvent{shape.HandleCreationCommand(command)}, nil
}

func (a StandardCommandApplier) ApplyStretchCommand(command StretchCommand) ([]cqrs.DomainEvent, error) {
	shape, err := a.loadShapeFromEventStore(command.id)
	if err != nil {
		return nil, err
	}

	return []cqrs.DomainEvent{shape.HandleStretchCommand(command)}, nil

}

func (a StandardCommandApplier) loadShapeFromEventStore(uuid uuid.UUID) (Shape, error) {
	_, events, err := a.eventStore.Load(context.TODO(), uuid)
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

func (a StandardCommandApplier) createShape(createdEvent cqrs.DomainEvent) Shape {
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
