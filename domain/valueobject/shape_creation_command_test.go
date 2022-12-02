package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullShapeCommand(t *testing.T) {
	shape := MockShape{}
	command, err := NewCreationShapeCommand(&shape)
	assert.NoError(t, err)
	assert.Equal(t, &shape, command.(newShapeCommand).shape)
}

func TestFullShapeCommandErrorWhenNoShapeProvided(t *testing.T) {
	_, err := NewCreationShapeCommand(nil)
	assert.Error(t, err)
}
