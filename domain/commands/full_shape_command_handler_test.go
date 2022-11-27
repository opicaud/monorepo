package commands

import (
	"example2/domain/utils"
	"example2/domain/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerACommand(t *testing.T) {
	fakeRepository := utils.FakeRepository{}
	command, _ := newFullShapeCommand(newFakeShape())
	handler := NewFullShapeCommandHandler(&fakeRepository)

	err := handler.Handle(*command)

	assert.NoError(t, err)
	fakeRepository.AssertContains(t, command.shape)
}

func TestAStandardHandlerACommand(t *testing.T) {
	fakeRepository := utils.NewFakeRepository()
	handler := NewFullShapeCommandHandler(fakeRepository)
	assert.IsType(t, &utils.FakeRepository{}, handler.(*fullShapeCommandHandler).repository)
}

type fakeShape struct{}

func (f *fakeShape) Area() error {
	return nil
}

func newFakeShape() valueobject.Shape {
	return new(fakeShape)
}
