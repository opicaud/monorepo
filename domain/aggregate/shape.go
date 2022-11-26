package aggregate

type Shape interface {
	Area() error
}

type rectangle struct {
	length float32
	width  float32
}

func (c *rectangle) Area() error {
	return nil
}
