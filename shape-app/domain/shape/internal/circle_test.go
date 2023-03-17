package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetAreaOfACircle(t *testing.T) {
	shape := circle{radius: 2}
	shape.calculateArea()
	assert.Equal(t, shape.area, float32(12.566371))
}
