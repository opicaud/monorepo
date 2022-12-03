package valueobject

import (
	"errors"
	"fmt"
)

type ShapeBuilder struct {
	nature    string
	builderOf map[string]func([]float32) Shape
}

type IShapeBuilder interface {
	CreateAShape(nature string) IShapeBuilder
	WithDimensions(dimensions []float32) (Shape, error)
}

func (s *ShapeBuilder) WithDimensions(dimensions []float32) (Shape, error) {
	builderOf := s.builderOf[s.nature]
	if builderOf == nil {
		return nil, errors.New(fmt.Sprintf("unable to create %s, this shape is unknown", s.nature))
	}
	return builderOf(dimensions), nil
}

func (s *ShapeBuilder) CreateAShape(nature string) IShapeBuilder {
	s.nature = nature
	return s
}

func newShapeBuilder() *ShapeBuilder {
	s := new(ShapeBuilder)
	s.builderOf = map[string]func(f []float32) Shape{
		"rectangle": func(f []float32) Shape { return newRectangle(f[0], f[1]) },
		"circle":    func(f []float32) Shape { return newCircle(f[0]) },
	}
	return s
}
