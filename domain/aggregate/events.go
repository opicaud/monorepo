package aggregate

import "github.com/google/uuid"

type ShapeCreatedEvent struct {
	nature     string
	dimensions []float32
	id         uuid.UUID
}

func (s ShapeCreatedEvent) AggregateId() uuid.UUID {
	return s.id
}

type AreaShapeCalculated struct {
	Area float32
	id   uuid.UUID
}

func (a AreaShapeCalculated) AggregateId() uuid.UUID {
	return a.id
}
