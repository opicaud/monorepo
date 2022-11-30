package valueobject

import (
	"errors"
	"fmt"
)

type ShapeBuilder struct {
	nature string
}

type IShapeBuilder interface {
	CreateAShape(nature string) IShapeBuilder
	WithDimensions(dimensions []float32) (Shape, error)
}

func (f *ShapeBuilder) WithDimensions(dimensions []float32) (Shape, error) {
	if "rectangle" == f.nature {
		return newRectangle(dimensions[0], dimensions[1]), nil
	}
	if "circle" == f.nature {
		return newCircle(dimensions[0]), nil
	}
	return nil, errors.New(fmt.Sprintf("unable to create %s, this shape is unknown", f.nature))
}

func (f *ShapeBuilder) CreateAShape(nature string) IShapeBuilder {
	f.nature = nature
	return f
}

func NewShapeBuilder() *ShapeBuilder {
	return new(ShapeBuilder)
}
