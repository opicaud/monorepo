package fullshapecommand

import (
	"example2/domain/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerACommand(t *testing.T) {
	fakeRepository := utils.FakeRepository{}
	shape := utils.CreateAMockShape()
	command, _ := NewFullShapeCommand(&shape)
	handler := NewFullShapeCommandHandler(&fakeRepository)

	err := handler.Execute(command)

	assert.NoError(t, err)
	shape.Mock.AssertCalled(t, "Area")
	fakeRepository.AssertContains(t, command.shape)

}

func TestAStandardHandlerACommand(t *testing.T) {
	fakeRepository := utils.NewFakeRepository()
	handler := NewFullShapeCommandHandler(fakeRepository)
	assert.IsType(t, &utils.FakeRepository{}, handler.(*fullShapeCommandHandler).repository)
}
