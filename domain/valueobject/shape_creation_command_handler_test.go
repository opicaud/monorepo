package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerACommand(t *testing.T) {
	fakeRepository := FakeRepository{}
	shape := CreateAMockShape()
	command, _ := NewCreationShapeCommand(shape)
	handler := NewShapeCreationCommandHandler(&fakeRepository)

	err := handler.Execute(command.(newShapeCommand))

	assert.NoError(t, err)
	shape.Mock.AssertCalled(t, "calculateArea")
	fakeRepository.AssertContains(t, command.(newShapeCommand).shape)

}

func TestAStandardHandlerACommand(t *testing.T) {
	fakeRepository := NewFakeRepository()
	handler := NewShapeCreationCommandHandler(fakeRepository)
	assert.IsType(t, &FakeRepository{}, handler.(*shapeCommandHandler).repository)
}
