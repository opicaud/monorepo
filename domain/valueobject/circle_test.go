package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetAreaOfACircle(t *testing.T) {
	shape := circle{radius: 2}
	shape.Area()
	assert.Equal(t, shape.GetArea(), float32(6.2831855))
}
