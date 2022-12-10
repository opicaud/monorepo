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

type Streched struct {
	Area       float32
	dimensions []float32
	id         uuid.UUID
}

func (a Streched) AggregateId() uuid.UUID {
	return a.id
}

func (a Streched) ApplyOn(shape Shape) Shape {
	return shape.ApplyShapeStrechedEvent(a)
}
