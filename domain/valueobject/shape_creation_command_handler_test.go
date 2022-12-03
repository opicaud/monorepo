package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerACommand(t *testing.T) {
	fakeRepository := FakeRepository{}
	command, _ := newCreationShapeCommand("rectangle", []float32{1, 2})
	handler := NewShapeCreationCommandHandler(&fakeRepository)

	err := handler.Execute(command.(newShapeCommand))

	assert.NoError(t, err)

}

func TestAStandardHandlerACommand(t *testing.T) {
	fakeRepository := NewFakeRepository()
	handler := NewShapeCreationCommandHandler(fakeRepository)
	assert.IsType(t, &FakeRepository{}, handler.(*shapeCommandHandler).repository)
}
