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

func (r *rectangle) HandleNewShape(command newShapeCommand) ShapeCreated {
	r.length = command.dimensions[0]
	r.width = command.dimensions[1]
	r.calculateArea()
	return ShapeCreated{id: r.id, dimensions: []float32{r.length, r.width}, Nature: "rectangle", Area: r.area}
}

func (r *rectangle) HandleStretchCommand(command newStretchCommand) ShapeStreched {
	r.length = command.stretchBy * r.length
	r.width = command.stretchBy * r.width
	r.calculateArea()
	return ShapeStreched{
		dimensions: []float32{r.length, r.width},
		Area:       r.area,
		id:         r.id,
	}
}

func (r *rectangle) ApplyShapeCreatedEvent(shapeCreatedEvent ShapeCreated) Shape {
	r.length = shapeCreatedEvent.dimensions[0]
	r.width = shapeCreatedEvent.dimensions[1]
	r.area = shapeCreatedEvent.Area
	r.id = shapeCreatedEvent.id
	return r
}

func (r *rectangle) ApplyShapeStrechedEvent(shapeStreched ShapeStreched) Shape {
	r.length = shapeStreched.dimensions[0]
	r.width = shapeStreched.dimensions[1]
	r.area = shapeStreched.Area
	return r
}

func newRectangle(id uuid.UUID, length float32, width float32) *rectangle {
	return &rectangle{id, length, width, 0}
}

func newRectangleWithId(id uuid.UUID) *rectangle {
	return &rectangle{id, 0, 0, 0}
}
