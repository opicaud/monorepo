package aggregate

import "github.com/google/uuid"

type ShapeEvent interface {
	Apply(shape Shape) Shape
}

type ShapeCreatedEvent struct {
	Nature     string
	dimensions []float32
	id         uuid.UUID
}

func (s ShapeCreatedEvent) AggregateId() uuid.UUID {
	return s.id
}

func (s ShapeCreatedEvent) Apply(shape Shape) Shape {
	return shape.ApplyShapeCreatedEvent(s)
}

type AreaShapeCalculated struct {
	Area float32
	id   uuid.UUID
}

func (a AreaShapeCalculated) AggregateId() uuid.UUID {
	return a.id
}

func (a AreaShapeCalculated) Apply(shape Shape) Shape {
	return shape.ApplyAreaShapeCalculated(a)
}
