package aggregate

import (
	"errors"
	"example2/infra"
	"fmt"
	"github.com/google/uuid"
)

type ShapeBuilder struct {
	nature    string
	id        uuid.UUID
	builderOf map[string]func([]float32) Shape
}

func (s *ShapeBuilder) withDimensions(dimensions []float32) (Shape, infra.Event, error) {
	builderOf := s.builderOf[s.nature]
	if builderOf == nil {
		return nil, nil, errors.New(fmt.Sprintf("unable to create %s, this shape is unknown", s.nature))
	}
	s.id = uuid.New()
	shape := builderOf(dimensions)
	return shape, ShapeCreatedEvent{id: s.id, Nature: s.nature, dimensions: dimensions}, nil
}

func (s *ShapeBuilder) createAShape(nature string) *ShapeBuilder {
	s.nature = nature
	return s
}

func newShapeBuilder() *ShapeBuilder {
	s := new(ShapeBuilder)
	s.builderOf = map[string]func(f []float32) Shape{
		"rectangle": func(f []float32) Shape { return newRectangle(s.id, f[0], f[1]) },
		"circle":    func(f []float32) Shape { return newCircle(s.id, f[0]) },
	}
	return s
}
