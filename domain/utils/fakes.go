package utils

import "example2/domain/valueobject"

type FakeShapeBuilder struct {
	Mock MockShape
}

func (s *FakeShapeBuilder) CreateAShape(nature string) valueobject.IShapeBuilder {
	return s
}

func (s *FakeShapeBuilder) WithDimensions(dimensions []float32) (valueobject.Shape, error) {
	s.Mock = CreateAMockShape()
	return &s.Mock, nil
}
