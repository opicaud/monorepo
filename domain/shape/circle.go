package shape

import (
	"github.com/google/uuid"
	"math"
)

type circle struct {
	id     uuid.UUID
	radius float32
	area   float32
}

func (r *circle) calculateArea() {
	r.area = r.radius * r.radius * math.Pi
}

func (r *circle) HandleNewShape(command newShapeCommand) Created {
	r.radius = command.dimensions[0]
	r.calculateArea()
	return newEventFactory().newShapeCreatedEvent(r.id, "circle", r.area, r.radius)

}

func (r *circle) HandleStretchCommand(command newStretchCommand) Streched {
	r.radius = command.stretchBy * r.radius
	r.calculateArea()
	return newEventFactory().newShapeStretchedEvent(r.id, r.area, r.radius)

}

func (r *circle) ApplyShapeCreatedEvent(shapeCreatedEvent Created) Shape {
	r.radius = shapeCreatedEvent.dimensions[0]
	r.area = shapeCreatedEvent.Area
	r.id = shapeCreatedEvent.id
	return r
}

func (r *circle) ApplyShapeStrechedEvent(shapeStreched Streched) Shape {
	r.radius = shapeStreched.dimensions[0]
	r.area = shapeStreched.Area
	return r
}

func newCircle(id uuid.UUID) *circle {
	return &circle{id, 0, 0}
}
