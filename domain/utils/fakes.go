package utils

import (
	"example2/domain/repository"
	"example2/domain/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func NewFakeRepository() repository.Repository {
	fakeRepository := new(FakeRepository)
	return fakeRepository
}

func (f *FakeRepository) Save(shape valueobject.Shape) error {
	f.Shapes = append(f.Shapes, shape)
	return nil
}

type FakeRepository struct {
	Shapes []valueobject.Shape
}

func (f *FakeRepository) AssertContains(t *testing.T, shape valueobject.Shape) bool {
	return assert.Contains(t, f.Shapes, shape)
}
