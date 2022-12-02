package valueobject

type Shape interface {
	HandleNewShape(command newShapeCommand)
	GetArea() float32
}
