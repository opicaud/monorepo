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

func (r *rectangle) HandleNewShape(command newShapeCommand) Created {
	r.length = command.dimensions[0]
	r.width = command.dimensions[1]
	r.calculateArea()
	return newEventFactory().newShapeCreatedEvent(r.id, "rectangle", r.area, r.length, r.width)
}

func (r *rectangle) HandleStretchCommand(command newStretchCommand) Streched {
	r.length = command.stretchBy * r.length
	r.width = command.stretchBy * r.width
	r.calculateArea()
	return newEventFactory().newShapeStretchedEvent(r.id, r.area, r.length, r.width)

}

func (r *rectangle) ApplyShapeCreatedEvent(shapeCreatedEvent Created) Shape {
	r.length = shapeCreatedEvent.dimensions[0]
	r.width = shapeCreatedEvent.dimensions[1]
	r.area = shapeCreatedEvent.Area
	r.id = shapeCreatedEvent.id
	return r
}

func (r *rectangle) ApplyShapeStrechedEvent(shapeStreched Streched) Shape {
	r.length = shapeStreched.dimensions[0]
	r.width = shapeStreched.dimensions[1]
	r.area = shapeStreched.Area
	return r
}

func newRectangle(id uuid.UUID) *rectangle {
	return &rectangle{id, 0, 0, 0}
}
