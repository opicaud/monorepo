package internal

import (
	"github.com/google/uuid"
)

type rectangle struct {
	id     uuid.UUID
	length float32
	width  float32
	area   float32
}

func (r *rectangle) calculateArea() {
	r.area = r.length * r.width
}

func (r *rectangle) HandleCreationCommand(command CreationCommand) Created {
	r.length = command.dimensions[0]
	r.width = command.dimensions[1]
	r.calculateArea()
	return NewShapeEventFactory().NewShapeCreatedEvent(r.id, "rectangle", r.area, r.length, r.width)
}

func (r *rectangle) HandleStretchCommand(command StretchCommand) Stretched {
	r.length = command.stretchBy * r.length
	r.width = command.stretchBy * r.width
	r.calculateArea()
	return NewShapeEventFactory().NewShapeStretchedEvent(r.id, r.area, r.length, r.width)

}

func (r *rectangle) ApplyCreatedEvent(shapeCreatedEvent Created) Shape {
	r.length = shapeCreatedEvent.Dimensions[0]
	r.width = shapeCreatedEvent.Dimensions[1]
	r.area = shapeCreatedEvent.Area
	r.id = shapeCreatedEvent.AggregateId()
	return r
}

func (r *rectangle) ApplyStretchedEvent(shapeStretched Stretched) Shape {
	r.length = shapeStretched.Dimensions[0]
	r.width = shapeStretched.Dimensions[1]
	r.area = shapeStretched.Area
	return r
}

func newRectangle(id uuid.UUID) *rectangle {
	return &rectangle{id, 0, 0, 0}
}
