package aggregate

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

func (r *circle) HandleNewShape(command newShapeCommand) AreaShapeCalculated {
	r.calculateArea()
	return AreaShapeCalculated{
		id:   r.id,
		Area: r.area,
	}
}

func (r *circle) HandleStretchCommand(command newStretchCommand) AreaShapeCalculated {
	r.radius = command.stretchBy * r.radius
	r.calculateArea()
	return AreaShapeCalculated{
		Area: r.area,
		id:   r.id,
	}
}

func (r *circle) ApplyShapeCreatedEvent(shapeCreatedEvent ShapeCreatedEvent) Shape {
	r.radius = shapeCreatedEvent.dimensions[0]
	return r
}

func (r *circle) ApplyAreaShapeCalculated(areaShapeCalculated AreaShapeCalculated) Shape {
	r.area = areaShapeCalculated.Area
	return r
}

func newCircle(id uuid.UUID, radius float32) *circle {
	return &circle{id, radius, 0}
}
