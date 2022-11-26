package valueobject

import "errors"

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
	return nil, errors.New("shape unknown")
}

func (f *ShapeBuilder) CreateAShape(nature string) IShapeBuilder {
	f.nature = nature
	return f
}

func newRectangle(length float32, width float32) *rectangle {
	return &rectangle{length, width}
}

func NewShapeBuilder() *ShapeBuilder {
	return new(ShapeBuilder)
}
