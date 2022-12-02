package shapecreationcommand

import (
	"example2/domain/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerACommand(t *testing.T) {
	fakeRepository := utils.FakeRepository{}
	shape := utils.CreateAMockShape()
	command, _ := NewCreationShapeCommand(shape)
	handler := NewShapeCreationCommandHandler(&fakeRepository)

	err := handler.Execute(command.(newShapeCommand))

	assert.NoError(t, err)
	shape.Mock.AssertCalled(t, "CalculateArea")
	fakeRepository.AssertContains(t, command.(newShapeCommand).shape)

}

func TestAStandardHandlerACommand(t *testing.T) {
	fakeRepository := utils.NewFakeRepository()
	handler := NewShapeCreationCommandHandler(fakeRepository)
	assert.IsType(t, &utils.FakeRepository{}, handler.(*shapeCommandHandler).repository)
}
