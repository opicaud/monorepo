package utils

import (
	"github.com/stretchr/testify/mock"
)

type MockShape struct {
	mock.Mock
}

func CreateAMockShape() *MockShape {
	shapeMock := MockShape{}
	shapeMock.On("Area").Return(nil)
	return &shapeMock
}

func (c *MockShape) Area() {
	c.Called()
}

func (c *MockShape) GetArea() float32 {
	return 0
}
