package valueobject

import (
	"example2/domain/commands"
	"math"
)

type circle struct {
	radius float32
	area   float32
}

func (r *circle) Area() {
	r.area = r.radius * math.Pi
}

func (r *circle) Execute(command commands.Command) error {
	r.Area()
	return nil
}

func (r circle) GetArea() float32 {
	return r.area
}

func newCircle(radius float32) *circle {
	return &circle{radius, 0}
}
