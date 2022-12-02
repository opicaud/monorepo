package utils

import (
	"example2/domain/commands"
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

func (r *MockShape) Execute(command commands.Command) error {
	r.Area()
	return nil
}

func (c *MockShape) GetArea() float32 {
	return 0
}
