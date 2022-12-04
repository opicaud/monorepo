package valueobject

import (
	"example2/infra"
	"math"
)

type circle struct {
	radius float32
	area   float32
}

func (r *circle) calculateArea() {
	r.area = r.radius * math.Pi
}

func (r *circle) HandleNewShape(command newShapeCommand) infra.Event {
	r.calculateArea()
	return AreaShapeCalculated{
		Area: r.area,
	}
}

func newCircle(radius float32) *circle {
	return &circle{radius, 0}
}
