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

func (r *rectangle) HandleCaculateShapeArea(command newShapeCommand) AreaShapeCalculated {
	r.calculateArea()
	return AreaShapeCalculated{
		Area: r.area,
		id:   r.id,
	}
}

func newRectangle(id uuid.UUID, length float32, width float32) *rectangle {
	return &rectangle{id, length, width, 0}
}
