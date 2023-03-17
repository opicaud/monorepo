package internal

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullShapeCommand(t *testing.T) {
	command := NewCreationShapeCommand("rectangle", []float32{1, 3})
	assert.NotNil(t, command)
}

func TestStretchShapeCommand(t *testing.T) {
	id := uuid.New()
	command := NewStretchShapeCommand(id, 1)
	assert.Equal(t, id, command.id)
	assert.Equal(t, float32(1), command.stretchBy)
}
