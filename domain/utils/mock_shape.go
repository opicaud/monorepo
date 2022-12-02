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
	shapeMock.On("CalculateArea").Return(nil)
	return &shapeMock
}

func (c *MockShape) CalculateArea() {
	c.Called()
}

func (r *MockShape) Execute(command commands.Command) error {
	r.CalculateArea()
	return nil
}

func (c *MockShape) GetArea() float32 {
	return 0
}
