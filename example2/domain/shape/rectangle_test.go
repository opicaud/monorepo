package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetAreaOfARectangle(t *testing.T) {
	shape := rectangle{length: 2, width: 3}
	shape.calculateArea()
	assert.Equal(t, shape.area, float32(6))
}
