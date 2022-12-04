package aggregate

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullShapeCommand(t *testing.T) {
	command, _ := newCreationShapeCommand("rectangle", []float32{1, 3})
	assert.NotNil(t, command)
}

func TestStretchShapeCommand(t *testing.T) {
	id := uuid.New()
	command := newStrechShapeCommand(id, 1)
	assert.Equal(t, id, command.(newStretchCommand).id)
	assert.Equal(t, float32(1), command.(newStretchCommand).stretchBy)

}
