package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetAreaOfARectangle(t *testing.T) {
	shape := rectangle{length: 2, width: 3}
	shape.Area()
	assert.Equal(t, shape.GetArea(), float32(6))
}
