package aggregate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewInMemoryRepository() *InMemoryRepository {
	fakeRepository := new(InMemoryRepository)
	return fakeRepository
}

func (f *InMemoryRepository) Save(shape Shape) error {
	if shape == nil {
		panic("shape is null")
	}
	f.Shapes = append(f.Shapes, shape)
	return nil
}

type InMemoryRepository struct {
	Shapes []Shape
}

func (f *InMemoryRepository) AssertContains(t *testing.T, shape Shape) bool {
	return assert.Contains(t, f.Shapes, shape)
}

func (f *InMemoryRepository) Get(i int) Shape {
	return f.Shapes[i]
}
