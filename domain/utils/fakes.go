package utils

import "example2/domain/aggregate"

type FakeShapeBuilder struct {
	Mock MockShape
}

func (s *FakeShapeBuilder) CreateAShape(nature string) aggregate.IShapeBuilder {
	return s
}

func (s *FakeShapeBuilder) WithDimensions(dimensions []float32) (aggregate.Shape, error) {
	s.Mock = CreateAMockShape()
	return &s.Mock, nil
}
