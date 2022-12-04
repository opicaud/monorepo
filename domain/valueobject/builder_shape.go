package valueobject

import (
	"errors"
	"fmt"
)

type ShapeBuilder struct {
	nature    string
	builderOf map[string]func([]float32) Shape
}

func (s *ShapeBuilder) withDimensions(dimensions []float32) (Shape, Event, error) {
	builderOf := s.builderOf[s.nature]
	if builderOf == nil {
		return nil, nil, errors.New(fmt.Sprintf("unable to create %s, this shape is unknown", s.nature))
	}
	return builderOf(dimensions), ShapeCreatedEvent{nature: s.nature, dimensions: dimensions}, nil
}

func (s *ShapeBuilder) createAShape(nature string) *ShapeBuilder {
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

type Event interface{}

type ShapeCreatedEvent struct {
	nature     string
	dimensions []float32
}

func (s ShapeCreatedEvent) Plop() Event {
	return s
}
