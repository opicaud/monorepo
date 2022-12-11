package shape

import "github.com/google/uuid"

type Event interface {
	ApplyOn(shape Shape) Shape
}

type Created struct {
	Nature     string
	dimensions []float32
	id         uuid.UUID
	Area       float32
}

func (s Created) AggregateId() uuid.UUID {
	return s.id
}

func (s Created) ApplyOn(shape Shape) Shape {
	return shape.ApplyShapeCreatedEvent(s)
}

type Stretched struct {
	Area       float32
	dimensions []float32
	id         uuid.UUID
}

func (a Stretched) AggregateId() uuid.UUID {
	return a.id
}

func (a Stretched) ApplyOn(shape Shape) Shape {
	return shape.ApplyShapeStretchedEvent(a)
}

type factoryEvents struct{}

func newEventFactory() *factoryEvents {
	return new(factoryEvents)
}

func (f factoryEvents) newShapeCreatedEvent(id uuid.UUID, nature string, area float32, dimensions ...float32) Created {
	return Created{id: id, dimensions: dimensions, Area: area, Nature: nature}
}

func (f factoryEvents) newShapeStretchedEvent(id uuid.UUID, area float32, dimensions ...float32) Stretched {
	return Stretched{id: id, dimensions: dimensions, Area: area}
}
