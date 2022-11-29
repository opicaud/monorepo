package valueobject

type rectangle struct {
	length float32
	width  float32
	area   float32
}

func (r *rectangle) Area() error {
	r.area = r.length * r.width
	return nil
}

func (r rectangle) GetArea() float32 {
	return r.area
}
