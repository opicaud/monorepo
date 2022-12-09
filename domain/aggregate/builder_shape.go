package aggregate

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type ShapeBuilder struct {
	nature            string
	id                uuid.UUID
	builderOf         map[string]func([]float32) Shape
	builderFromNature map[string]func() Shape
}

func (s *ShapeBuilder) createAShape(nature string) *ShapeBuilder {
	s.nature = nature
	return s
}

func (s *ShapeBuilder) withId(id uuid.UUID) (Shape, error) {
	s.id = id
	builderOf := s.builderFromNature[s.nature]
	if builderOf == nil {
		return nil, errors.New(fmt.Sprintf("unable to create %s, this shape is unknown", s.nature))
	}
	shape := builderOf()
	return shape, nil
}

func newShapeBuilder() *ShapeBuilder {
	s := new(ShapeBuilder)
	s.id = uuid.New()
	s.builderOf = map[string]func(f []float32) Shape{
		"rectangle": func(f []float32) Shape { return newRectangle(s.id, f[0], f[1]) },
		"circle":    func(f []float32) Shape { return newCircle(s.id, f[0]) },
	}
	s.builderFromNature = map[string]func() Shape{
		"rectangle": func() Shape { return newRectangleWithId(s.id) },
		"circle":    func() Shape { return newCircleWithId(s.id) },
	}
	return s
}
