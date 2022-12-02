package valueobject

import (
	"github.com/stretchr/testify/mock"
)

type MockShape struct {
	mock.Mock
}

func CreateAMockShape() *MockShape {
	shapeMock := MockShape{}
	shapeMock.On("calculateArea").Return(nil)
	return &shapeMock
}

func (c *MockShape) calculateArea() {
	c.Called()
}

func (c *MockShape) GetArea() float32 {
	return 0
}

func (r *MockShape) HandleNewShape(command newShapeCommand) {
	r.calculateArea()
}
