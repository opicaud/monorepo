package valueobject

type rectangle struct {
	length float32
	width  float32
	area   float32
}

func (r *rectangle) Area() {
	r.area = r.length * r.width
}

func (r rectangle) GetArea() float32 {
	return r.area
}
