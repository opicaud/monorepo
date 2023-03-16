package shape

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
)

type Event interface {
	ApplyOn(shape Shape) Shape
}

type Created struct {
	Nature     string
	Dimensions []float32
	id         uuid.UUID
	Area       float32
}

func (s Created) AggregateId() uuid.UUID {
	return s.id
}

func (s Created) Name() string {
	return "SHAPE_CREATED"
}

func (s Created) Data() []byte {
	marshal, _ := json.Marshal(s)
	return marshal
}

func (s Created) ApplyOn(shape Shape) Shape {
	return shape.ApplyShapeCreatedEvent(s)
}

type Stretched struct {
	Area       float32
	Dimensions []float32
	id         uuid.UUID
}

func (a Stretched) AggregateId() uuid.UUID {
	return a.id
}

func (a Stretched) Name() string {
	return "SHAPE_STRETCHED"
}

func (a Stretched) Data() []byte {
	marshal, _ := json.Marshal(a)
	return marshal
}

func (a Stretched) ApplyOn(shape Shape) Shape {
	return shape.ApplyShapeStretchedEvent(a)
}

type factoryEvents struct{}

func newEventFactory() *factoryEvents {
	return new(factoryEvents)
}

func (f factoryEvents) newShapeCreatedEvent(id uuid.UUID, nature string, area float32, dimensions ...float32) Created {
	return Created{id: id, Dimensions: dimensions, Area: area, Nature: nature}
}

func (f factoryEvents) newShapeStretchedEvent(id uuid.UUID, area float32, dimensions ...float32) Stretched {
	return Stretched{id: id, Dimensions: dimensions, Area: area}
}

func (f factoryEvents) newDeserializedEvent(aggregateId uuid.UUID, event pkg.DomainEvent) pkg.DomainEvent {
	switch event.Name() {
	case "SHAPE_CREATED":
		v := &Created{}
		_ = json.Unmarshal(event.Data(), v)
		v.id = aggregateId
		return v
	case "SHAPE_STRETCHED":
		v := &Stretched{}
		_ = json.Unmarshal(event.Data(), v)
		v.id = aggregateId
		return v
	default:
		panic(fmt.Errorf("%s is not known as event", event.Name()))
	}

}
