package valueobject

import "math"

type circle struct {
	radius float32
	area   float32
}

func (r *circle) Area() {
	r.area = r.radius * math.Pi
}

func (r circle) GetArea() float32 {
	return r.area
}

func newCircle(radius float32) *circle {
	return &circle{radius, 0}
}
