package valueobject

import (
	"example2/infra"
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

func (r *rectangle) HandleNewShape(command newShapeCommand) infra.Event {
	r.calculateArea()
	return AreaShapeCalculated{
		area: r.area,
	}
}

func (r rectangle) GetArea() float32 {
	return r.area
}

func newRectangle(length float32, width float32) *rectangle {
	return &rectangle{uuid.New(), length, width, 0}
}

type AreaShapeCalculated struct {
	area float32
}
