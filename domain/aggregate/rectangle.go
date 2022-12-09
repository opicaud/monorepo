package aggregate

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

func (r *rectangle) HandleNewShape(command newShapeCommand) AreaShapeCalculated {
	r.calculateArea()
	return AreaShapeCalculated{
		Area: r.area,
		id:   r.id,
	}
}

func (r *rectangle) HandleStretchCommand(command newStretchCommand) AreaShapeCalculated {
	r.length = command.stretchBy * r.length
	r.width = command.stretchBy * r.width
	r.calculateArea()
	return AreaShapeCalculated{
		Area: r.area,
		id:   r.id,
	}
}

func (r *rectangle) ApplyShapeCreatedEvent(shapeCreatedEvent ShapeCreatedEvent) Shape {
	r.length = shapeCreatedEvent.dimensions[0]
	r.width = shapeCreatedEvent.dimensions[1]
	return r
}

func (r *rectangle) ApplyAreaShapeCalculated(areaShapeCalculated AreaShapeCalculated) Shape {
	r.area = areaShapeCalculated.Area
	return r
}

func newRectangle(id uuid.UUID, length float32, width float32) *rectangle {
	return &rectangle{id, length, width, 0}
}
