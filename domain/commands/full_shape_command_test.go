package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullShapeCommand(t *testing.T) {
	command, err := newFullShapeCommand("nature", 1, 2)
	assert.NoError(t, err)
	assert.Equal(t, command.nature, "nature")
	assert.Equal(t, command.dimensions, []float32{1, 2})
}

func TestFullShapeCommandErrorWhenNoDimensionsProvided(t *testing.T) {
	_, err := newFullShapeCommand("nature")
	assert.Error(t, err)
}
