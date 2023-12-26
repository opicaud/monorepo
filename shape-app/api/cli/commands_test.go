package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDefaultCommand(t *testing.T) {
	dimensions := []float32{3, 4}
	shape := "rectangle"
	shapes := CreateDefaultShapeRequest().GetShapes()
	assert.Equal(t, shapes.Shape, shape)
	assert.Equal(t, shapes.Dimensions, dimensions)
}

func TestCreateCommand(t *testing.T) {
	dimensions := []float32{3}
	shape := "circle"
	shapes := CreateShapeRequest(shape, dimensions).GetShapes()
	assert.Equal(t, shapes.Shape, shape)
	assert.Equal(t, shapes.Dimensions, dimensions)
}
