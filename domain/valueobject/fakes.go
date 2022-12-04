package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewFakeRepository() *FakeRepository {
	fakeRepository := new(FakeRepository)
	return fakeRepository
}

func (f *FakeRepository) Save(shape Shape) error {
	f.Shapes = append(f.Shapes, shape)
	return nil
}

type FakeRepository struct {
	Shapes []Shape
}

func (f *FakeRepository) AssertContains(t *testing.T, shape Shape) bool {
	return assert.Contains(t, f.Shapes, shape)
}

func (f *FakeRepository) Get(i int) Shape {
	return f.Shapes[i]
}
