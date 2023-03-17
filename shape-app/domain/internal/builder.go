package internal

import (
	"fmt"
	"github.com/google/uuid"
)

type Builder struct {
	nature            string
	id                uuid.UUID
	builderFromNature map[string]func() Shape
}

func (s *Builder) withNature(nature string) *Builder {
	s.nature = nature
	return s
}

func (s *Builder) withId(id uuid.UUID) (Shape, error) {
	s.id = id
	builderOf := s.builderFromNature[s.nature]
	if builderOf == nil {
		return nil, fmt.Errorf("unable to create %s, this shape is unknown", s.nature)
	}
	shape := builderOf()
	return shape, nil
}

func newShapeBuilder() *Builder {
	s := new(Builder)
	s.id = uuid.New()
	s.builderFromNature = map[string]func() Shape{
		"rectangle": func() Shape { return newRectangle(s.id) },
		"circle":    func() Shape { return newCircle(s.id) },
	}
	return s
}
