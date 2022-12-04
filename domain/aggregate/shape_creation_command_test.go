package aggregate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullShapeCommand(t *testing.T) {
	command := createCommand("rectangle", []float32{1, 3})
	assert.NotNil(t, command)
}
