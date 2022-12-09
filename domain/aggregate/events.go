package aggregate

import "github.com/google/uuid"

type ShapeEvent interface {
	Apply(shape Shape) Shape
}

type ShapeCreated struct {
	Nature     string
	dimensions []float32
	id         uuid.UUID
	Area       float32
}

func (s ShapeCreated) AggregateId() uuid.UUID {
	return s.id
}

func (s ShapeCreated) Apply(shape Shape) Shape {
	return shape.ApplyShapeCreatedEvent(s)
}

type ShapeStreched struct {
	Area       float32
	dimensions []float32
	id         uuid.UUID
}

func (a ShapeStreched) AggregateId() uuid.UUID {
	return a.id
}

func (a ShapeStreched) Apply(shape Shape) Shape {
	return shape.ApplyShapeStrechedEvent(a)
}
