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

func (r circle) GetArea() float32 {
	return r.area
}

func (r *circle) HandleNewShape(command newShapeCommand) infra.Event {
	r.calculateArea()
	return nil
}

func newCircle(radius float32) *circle {
	return &circle{radius, 0}
}
