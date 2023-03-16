package shape

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

func (r *rectangle) HandleNewShape(command CreationCommand) Created {
	r.length = command.dimensions[0]
	r.width = command.dimensions[1]
	r.calculateArea()
	return newShapeEventFactory().newShapeCreatedEvent(r.id, "rectangle", r.area, r.length, r.width)
}

func (r *rectangle) HandleStretchCommand(command StretchCommand) Stretched {
	r.length = command.stretchBy * r.length
	r.width = command.stretchBy * r.width
	r.calculateArea()
	return newShapeEventFactory().newShapeStretchedEvent(r.id, r.area, r.length, r.width)

}

func (r *rectangle) ApplyShapeCreatedEvent(shapeCreatedEvent Created) Shape {
	r.length = shapeCreatedEvent.Dimensions[0]
	r.width = shapeCreatedEvent.Dimensions[1]
	r.area = shapeCreatedEvent.Area
	r.id = shapeCreatedEvent.AggregateId()
	return r
}

func (r *rectangle) ApplyShapeStretchedEvent(shapeStretched Stretched) Shape {
	r.length = shapeStretched.Dimensions[0]
	r.width = shapeStretched.Dimensions[1]
	r.area = shapeStretched.Area
	return r
}

func newRectangle(id uuid.UUID) *rectangle {
	return &rectangle{id, 0, 0, 0}
}
