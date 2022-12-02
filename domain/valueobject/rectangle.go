package valueobject

import "example2/domain/commands"

type rectangle struct {
	length float32
	width  float32
	area   float32
}

func (r *rectangle) Area() {
	r.area = r.length * r.width
}

func (r *rectangle) Execute(command commands.Command) error {
	r.Area()
	return nil
}

func (r rectangle) GetArea() float32 {
	return r.area
}

func newRectangle(length float32, width float32) *rectangle {
	return &rectangle{length, width, 0}
}
