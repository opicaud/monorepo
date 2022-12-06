package aggregate

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type ShapeBuilder struct {
	nature    string
	id        uuid.UUID
	builderOf map[string]func([]float32) Shape
}

func (s *ShapeBuilder) withDimensions(dimensions []float32) (Shape, ShapeCreatedEvent, error) {
	builderOf := s.builderOf[s.nature]
	if builderOf == nil {
		return nil, ShapeCreatedEvent{}, errors.New(fmt.Sprintf("unable to create %s, this shape is unknown", s.nature))
	}
	shape := builderOf(dimensions)
	return shape, ShapeCreatedEvent{id: s.id, Nature: s.nature, dimensions: dimensions}, nil
}

func (s *ShapeBuilder) createAShape(nature string) *ShapeBuilder {
	s.nature = nature
	return s
}

func (s *ShapeBuilder) withId(id uuid.UUID) *ShapeBuilder {
	s.id = id
	return s
}

func newShapeBuilder() *ShapeBuilder {
	s := new(ShapeBuilder)
	s.id = uuid.New()
	s.builderOf = map[string]func(f []float32) Shape{
		"rectangle": func(f []float32) Shape { return newRectangle(s.id, f[0], f[1]) },
		"circle":    func(f []float32) Shape { return newCircle(s.id, f[0]) },
	}
	return s
}
