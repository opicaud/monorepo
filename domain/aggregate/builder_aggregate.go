package aggregate

import "errors"

type factoryOfShape struct {
	nature string
}

func (f *factoryOfShape) CreateAShape(nature string) *factoryOfShape {
	f.nature = nature
	return f
}

func (f *factoryOfShape) WithDimensions(dimensions ...float32) (Shape, error) {
	if "rectangle" == f.nature {
		return newRectangle(0, 0), nil

	}
	return nil, errors.New("shape unknown")
}

func newRectangle(length float32, width float32) *rectangle {
	return &rectangle{length, width}
}

func NewShapeBuilder() *factoryOfShape {
	return new(factoryOfShape)
}
