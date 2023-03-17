package internal

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

func (r *circle) HandleCreationCommand(command CreationCommand) Created {
	r.radius = command.dimensions[0]
	r.calculateArea()
	return NewShapeEventFactory().NewShapeCreatedEvent(r.id, "circle", r.area, r.radius)

}

func (r *circle) HandleStretchCommand(command StretchCommand) Stretched {
	r.radius = command.stretchBy * r.radius
	r.calculateArea()
	return NewShapeEventFactory().NewShapeStretchedEvent(r.id, r.area, r.radius)

}

func (r *circle) ApplyCreatedEvent(shapeCreatedEvent Created) Shape {
	r.radius = shapeCreatedEvent.Dimensions[0]
	r.area = shapeCreatedEvent.Area
	r.id = shapeCreatedEvent.Id
	return r
}

func (r *circle) ApplyStretchedEvent(shapeStretched Stretched) Shape {
	r.radius = shapeStretched.Dimensions[0]
	r.area = shapeStretched.Area
	return r
}

func newCircle(id uuid.UUID) *circle {
	return &circle{id, 0, 0}
}
