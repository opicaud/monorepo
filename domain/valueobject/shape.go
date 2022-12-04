package valueobject

type Shape interface {
	HandleNewShape(command newShapeCommand) Event
	GetArea() float32
}
