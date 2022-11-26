package valueobject

type Shape interface {
	Area() error
}

type rectangle struct {
	length float32
	width  float32
}

func (r rectangle) Area() error {
	return nil
}
